package http

type BaseResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
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

func MapBaseResponseWithMeta(status, message string, data interface{}, meta interface{}) BaseResponseWithMeta {
	return BaseResponseWithMeta{
		BaseResponse: MapBaseResponse(status, message, data),
		Meta:         meta,
	}
}
