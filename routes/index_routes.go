package routes

import "github.com/gorilla/mux"

// RegisterRoutes -> centraliza todas las rutas
func RegisterRoutes(r *mux.Router) {
	// Home
	RegisterHomeRoute(r)

	// Users
	RegisterUserRoutes(r)

	// Tasks
	RegisterTaskRoutes(r)
}
