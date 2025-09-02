package routes

import (
	"net/http"

	"github.com/F3AR-DEV/ApiRestGO/api/controllers"
	"github.com/F3AR-DEV/ApiRestGO/api/middlewares"
	"github.com/F3AR-DEV/ApiRestGO/core/dto"
	"github.com/gorilla/mux"
)

func ItemsRoutes(r *mux.Router) {
	items := r.PathPrefix("/items").Subrouter()

	// Aplicar middleware JWT a todas las rutas
	items.Use(middlewares.JWTMiddleware)

	// GET /items → listar items
	items.HandleFunc("", controllers.GetItems).Methods("GET")

	// GET /items/{id} → obtener item por id
	items.HandleFunc("/{id}", controllers.GetItemByID).Methods("GET")

	// POST /items → crear item con validación
	items.Handle("", middlewares.ValidateBodyMiddleware(
		http.HandlerFunc(controllers.CreateItem),
		dto.ItemRequest{},
	)).Methods("POST")

	// PUT /items/{id} → actualizar item con validación
	items.Handle("/{id}", middlewares.ValidateBodyMiddleware(
		http.HandlerFunc(controllers.UpdateItem),
		dto.ItemRequest{},
	)).Methods("PUT")

	// DELETE /items/{id} → eliminar item
	items.HandleFunc("/{id}", controllers.DeleteItem).Methods("DELETE")
}
