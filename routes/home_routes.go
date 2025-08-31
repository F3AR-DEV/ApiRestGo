package routes

import (
	"github.com/F3AR-DEV/ApiRestGO/controllers"
	"github.com/gorilla/mux"
)

func RegisterHomeRoute(r *mux.Router) {
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
}
