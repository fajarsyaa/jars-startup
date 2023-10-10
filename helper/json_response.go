package helper

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func JSONResponse(message, status string, code int, data interface{}) Response {
	meta := Meta{
		Status:  status,
		Message: message,
		Code:    code,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}
