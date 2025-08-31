package services

import (
	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/models"
)

// Obtener todas las tasks
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := db.DB.Find(&tasks)
	return tasks, result.Error
}

// Obtener task por ID
func GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	result := db.DB.First(&task, id)
	return task, result.Error
}

// Crear nueva task
func CreateTask(task *models.Task) error {
	result := db.DB.Create(task)
	return result.Error
}

// Eliminar task
func DeleteTask(id string) error {
	var task models.Task
	result := db.DB.First(&task, id)
	if task.ID == 0 {
		return result.Error
	}
	db.DB.Delete(&task)
	return nil
}
