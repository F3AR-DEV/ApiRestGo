package migrations

import (
	"log"

	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/data/models"
)

// RunMigrations ejecuta todas las migraciones de la DB
func RunMigrations() {
	err := db.DB.AutoMigrate(
		&models.Item{},
		// aquí puedes agregar más modelos: &models.User{}, &models.Task{}, etc.
	)

	if err != nil {
		log.Fatal("❌ Error en migraciones:", err)
	} else {
		log.Println("✅ Migraciones ejecutadas correctamente")
	}
}
