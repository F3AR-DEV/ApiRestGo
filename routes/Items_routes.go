package routes

import (
	"github.com/F3AR-DEV/ApiRestGO/controllers"
	"github.com/gorilla/mux"
)

func ItemsRoutes(r *mux.Router) {
	items := r.PathPrefix("/items").Subrouter()
	items.HandleFunc("", controllers.CreateItem).Methods("POST")
	items.HandleFunc("", controllers.GetItems).Methods("GET")
	items.HandleFunc("/{id}", controllers.GetItemByID).Methods("GET")
	items.HandleFunc("/{id}", controllers.UpdateItem).Methods("PUT")
	items.HandleFunc("/{id}", controllers.DeleteItem).Methods("DELETE")
}
