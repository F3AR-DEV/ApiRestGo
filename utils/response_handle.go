package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/F3AR-DEV/ApiRestGO/business/dto"
)

// NewAPIResponse construye un ApiResponse con fecha y hora actual
func NewAPIResponse(status string, message string, data interface{}) dto.ApiResponse {
	now := time.Now()
	return dto.ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Date:    now.Format("2006-01-02"),
		Time:    now.Format("15:04:05"),
	}
}

// WriteJSONResponse manda la respuesta al cliente
func WriteJSONResponse(w http.ResponseWriter, statusCode int, status string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := NewAPIResponse(status, message, data)

	_ = json.NewEncoder(w).Encode(resp)
}
