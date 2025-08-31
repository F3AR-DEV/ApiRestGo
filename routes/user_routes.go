package routes

import (
	"github.com/F3AR-DEV/ApiRestGO/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", controllers.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.DeleteUsersHandler).Methods("DELETE")
}
