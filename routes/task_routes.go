package routes

import (
	"github.com/F3AR-DEV/ApiRestGO/controllers"
	"github.com/F3AR-DEV/ApiRestGO/middleware"
	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(r *mux.Router) {
	tasks := r.PathPrefix("/tasks").Subrouter()
	tasks.Use(middleware.LoggerMiddleware) // se ejecuta primero
	//tasks.Use(middleware.AuthMiddleware)   // se ejecuta después del Logger

	tasks.HandleFunc("", controllers.GetTasksHandler).Methods("GET")
	tasks.HandleFunc("", controllers.PostTasksHandler).Methods("POST")
	tasks.HandleFunc("/{id}", controllers.GetTaskHandler).Methods("GET")
	tasks.HandleFunc("/{id}", controllers.DeleteTasksHandler).Methods("DELETE")
}
