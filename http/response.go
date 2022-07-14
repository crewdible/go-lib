package http

type BaseResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

func MapBaseResponse(status, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
