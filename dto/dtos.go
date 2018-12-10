package dto

type HelloWorldResponse struct {
	Message string `json:"message"`
}

func NewHelloWorldResponse(message string) *HelloWorldResponse {
	return &HelloWorldResponse{Message: message}
}
