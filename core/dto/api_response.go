package dto

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Date    string      `json:"date"`
	Time    string      `json:"time"`
}
