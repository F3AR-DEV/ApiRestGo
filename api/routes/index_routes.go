package routes

import "github.com/gorilla/mux"

// RegisterRoutes -> centraliza todas las rutas
func RegisterRoutes(r *mux.Router) {
	// Home
	RegisterHomeRoute(r)
	//Auth
	RegisterAuthRoute(r)
	//Items
	ItemsRoutes(r)

}
