package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// LoggerMiddleware ya no solo imprime, sino que puede bloquear si detecta algo
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		fmt.Println("[Logger] user-agent:", userAgent)

		// Ejemplo: bloquear si no viene user-agent
		if userAgent == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "User-Agent requerido"})
			return
		}

		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware verifica un token y retorna error JSON si no es válido
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer miTokenSecreto" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "No autorizado"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
