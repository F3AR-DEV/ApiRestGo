package services

import (
	"errors"
	"strconv"

	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/F3AR-DEV/ApiRestGO/repositories"

	"github.com/F3AR-DEV/ApiRestGO/config/db"
)

var taskRepo = &repositories.TaskRepository{
	DB: db.DB, // función que devuelve *gorm.DB ya inicializado
}

func GetAllTasks() ([]models.Task, error) {
	return taskRepo.FindAll()
}

func GetTaskByID(idStr string) (models.Task, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.Task{}, errors.New("ID inválido")
	}
	return taskRepo.FindByID(uint(id))
}

func CreateTask(task *models.Task) (*models.Task, error) {
	created, err := taskRepo.Save(*task)
	return &created, err
}

func DeleteTask(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("ID inválido")
	}
	return taskRepo.Delete(uint(id))
}
