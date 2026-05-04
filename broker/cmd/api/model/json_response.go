package model

type JsonResponse struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Data    any    `json:"data,omitempty"`
}
