package controller

type AppResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message"`
}
