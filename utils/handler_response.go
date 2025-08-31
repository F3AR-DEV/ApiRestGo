package utils

import (
	"encoding/json"
	"net/http"
)

// Estructura de respuesta genérica
type APIResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
	ErrorDetail string      `json:"error,omitempty"`
}

// Respuesta de éxito
func RespondSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Respuesta de error
func RespondError(w http.ResponseWriter, statusCode int, message string, errorDetail string) {
	response := APIResponse{
		Success:     false,
		Message:     message,
		ErrorDetail: errorDetail,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
