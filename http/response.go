package http

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
