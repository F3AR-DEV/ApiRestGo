package repositories

import (
	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/data/models"
)

// CreateUser guarda un nuevo usuario en la base de datos
func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

// GetUserByEmail obtiene un usuario por email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
