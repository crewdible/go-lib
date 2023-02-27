package http

import (
	"encoding/json"
	"net/http"

	"github.com/crewdible/go-lib/errors"
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type CrewBaseResponse struct {
	Result  int         `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type BaseResponseWithMeta struct {
	BaseResponse
	Meta interface{} `json:"meta"`
}

func MapBaseResponse(status, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func MapCrewBaseResponse(result int, message string, data interface{}) CrewBaseResponse {
	return CrewBaseResponse{
		Result:  result,
		Message: message,
		Data:    data,
	}
}

func MapBaseResponseWithMeta(status, message string, data interface{}, meta interface{}) BaseResponseWithMeta {
	return BaseResponseWithMeta{
		BaseResponse: MapBaseResponse(status, message, data),
		Meta:         meta,
	}
}

func RespondSuccessJSON(c echo.Context, data interface{}) error {
	resp := BaseResponse{
		Data:   data,
		Status: "success",
	}

	dataByte, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return err
	}

	c.Set("response", dataByte)

	if delay, ok := c.Get("delayResponse").(bool); ok && delay {
		return nil
	}

	return c.Blob(http.StatusOK, "application/json", dataByte)
}

func RespondErrorJSON(c echo.Context, err error) error {
	resp := BaseResponse{
		Status:  "failed",
		Message: err.Error(),
	}
	b, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return err
	}
	c.Set("response", b)

	return c.Blob(errors.GetErrorHttpCode(err), "application/json", b)
}
