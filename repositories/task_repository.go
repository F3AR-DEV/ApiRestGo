package repositories

import (
	"github.com/F3AR-DEV/ApiRestGO/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

// Guardar una nueva tarea
func (r *TaskRepository) Save(task models.Task) (models.Task, error) {
	if err := r.DB.Create(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

// Obtener todas las tareas
func (r *TaskRepository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Buscar una tarea por ID
func (r *TaskRepository) FindByID(id uint) (models.Task, error) {
	var task models.Task
	if err := r.DB.First(&task, id).Error; err != nil {
		return task, err
	}
	return task, nil
}

// Actualizar una tarea
func (r *TaskRepository) Update(task models.Task) (models.Task, error) {
	if err := r.DB.Save(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

// Eliminar una tarea por ID
func (r *TaskRepository) Delete(id uint) error {
	if err := r.DB.Delete(&models.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
