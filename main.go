package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/F3AR-DEV/ApiRestGO/api/routes"
	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/data/migrations"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Cargar .env
	_ = godotenv.Load()

	// 2. Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// 3. Conexión a BD
	db.DBConnection()
	fmt.Println("✅ Conexión lista a la base de datos")
	migrations.RunMigrations()
	// 4. Router
	r := mux.NewRouter()

	// 5. Registrar rutas (igual que app.use(router))
	routes.RegisterRoutes(r)

	// 6. CORS middleware
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// 7. Levantar servidor
	fmt.Printf("🚀 Servidor listo en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, cors(r)))
}
