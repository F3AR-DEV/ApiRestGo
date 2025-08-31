package controllers

import (
	"net/http"

	dto "github.com/F3AR-DEV/ApiRestGO/interfaces"
	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/F3AR-DEV/ApiRestGO/services"
	"github.com/F3AR-DEV/ApiRestGO/utils"
	"github.com/gorilla/mux"
)

// GET /tasks
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := services.GetAllTasks()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error al obtener tareas", err.Error())
		return
	}
	utils.RespondSuccess(w, http.StatusOK, "Tareas obtenidas", tasks)
}

// GET /tasks/{id}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task, err := services.GetTaskByID(params["id"])
	if err != nil || task.ID == 0 {
		utils.RespondError(w, http.StatusNotFound, "Tarea no encontrada", "")
		return
	}
	utils.RespondSuccess(w, http.StatusOK, "Tarea encontrada", task)
}

// POST /tasks
func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.TaskRequest

	// Decodificar y validar
	if err := utils.ParseAndValidate(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Datos inválidos", err.Error())
		return
	}

	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
		UserID:      req.UserID,
	}

	createdTask, err := services.CreateTask(&task)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "No se pudo crear la tarea", err.Error())
		return
	}

	utils.RespondSuccess(w, http.StatusCreated, "Tarea creada", createdTask)
}

// DELETE /tasks/{id}
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := services.DeleteTask(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Tarea no encontrada", err.Error())
		return
	}
	utils.RespondSuccess(w, http.StatusOK, "Tarea eliminada", nil)
}
