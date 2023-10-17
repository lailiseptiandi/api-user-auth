package utils

type apiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(data interface{}, message string) apiResponse {
	resp := apiResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return resp
}

func ResponseError(data interface{}, message string) apiResponse {
	resp := apiResponse{
		Status:  false,
		Message: message,
		Data:    data,
	}
	return resp
}
