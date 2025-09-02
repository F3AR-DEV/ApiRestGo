package services

import (
	"errors"

	"github.com/F3AR-DEV/ApiRestGO/data/models"
	"github.com/F3AR-DEV/ApiRestGO/data/repositories"
)

// CreateUser valida y crea un usuario
func CreateUser(user *models.User) error {
	if user.Email == "" {
		return errors.New("el email no puede estar vacío")
	}
	if user.Password == "" {
		return errors.New("la contraseña no puede estar vacía")
	}

	// Validación de email duplicado
	existingUser, _ := repositories.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("el email ya está en uso")
	}

	// Guardar usuario en DB
	return repositories.CreateUser(user)
}

// GetUserByEmail obtiene un usuario por email
func GetUserByEmail(email string) (*models.User, error) {
	return repositories.GetUserByEmail(email)
}
