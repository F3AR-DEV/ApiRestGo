package routes

import (
	"github.com/F3AR-DEV/ApiRestGO/controllers"
	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", controllers.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", controllers.PostTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", controllers.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTasksHandler).Methods("DELETE")
}
