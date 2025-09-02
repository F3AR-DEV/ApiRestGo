package routes

import (
	"github.com/F3AR-DEV/ApiRestGO/api/controllers"
	"github.com/F3AR-DEV/ApiRestGO/api/middlewares"
	"github.com/gorilla/mux"
)

func RegisterHomeRoute(r *mux.Router) {
	home := r.PathPrefix("/").Subrouter()
	home.Use(middlewares.LoggerMiddleware) // se ejecuta primero
	//home.Use(middleware.AuthMiddleware)   // se ejecuta después del Logger
	home.HandleFunc("/", controllers.HomeHandler).Methods("GET")
}
