package main

import (
	"net/http"

	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	/* 	db.DB.AutoMigrate(models.Task{})
	   	db.DB.AutoMigrate(models.User{}) */

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
