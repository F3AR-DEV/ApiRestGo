package services

import (
	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/models"
)

// Get all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	return users, result.Error
}

// Get user by ID
func GetUserByID(id string) (models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if user.ID != 0 {
		db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	}
	return user, result.Error
}

// Create a new user
func CreateUser(user *models.User) error {
	result := db.DB.Create(user)
	return result.Error
}

// Delete user
func DeleteUser(id string) error {
	var user models.User
	result := db.DB.First(&user, id)
	if user.ID == 0 {
		return result.Error
	}
	db.DB.Delete(&user)
	return nil
}
