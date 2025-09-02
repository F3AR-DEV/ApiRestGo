package routes

import (
	"net/http"

	"github.com/F3AR-DEV/ApiRestGO/api/controllers"
	"github.com/F3AR-DEV/ApiRestGO/api/middlewares"
	"github.com/F3AR-DEV/ApiRestGO/core/dto"

	"github.com/gorilla/mux"
)

func RegisterAuthRoute(r *mux.Router) {
	// POST /register → payload validado con RegisterRequest
	r.Handle("/register", middlewares.ValidateBodyMiddleware(
		http.HandlerFunc(controllers.RegisterHandler),
		dto.RegisterRequest{},
	)).Methods("POST")

	// POST /login → payload validado con LoginRequest
	r.Handle("/login", middlewares.ValidateBodyMiddleware(
		http.HandlerFunc(controllers.LoginHandler),
		dto.LoginRequest{},
	)).Methods("POST")
}
