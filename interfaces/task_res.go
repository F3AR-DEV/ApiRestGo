package dto

type TaskResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
	ErrorDetail string      `json:"error,omitempty"`
}
