package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	dto "github.com/F3AR-DEV/ApiRestGO/interfaces"
	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/F3AR-DEV/ApiRestGO/services"
	"github.com/F3AR-DEV/ApiRestGO/utils"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := services.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task, err := services.GetTaskByID(params["id"])
	if err != nil || task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task Not Found"))
		return
	}
	json.NewEncoder(w).Encode(task)
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
    var req dto.TaskRequest

    // Usar helper genérico (decodifica y valida)
    if err := utils.ParseAndValidate(r, &req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Mapear DTO -> Modelo
    task := models.Task{
        Title:       req.Title,
        Description: req.Description,
        Done:        req.Done,
        UserID:      req.UserID,
    }

	fmt.Println(reflect.TypeOf(task))

    // Guardar en la BD
    if err := services.CreateTask(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Respuesta (ya guardada con ID y timestamps)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}


func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := services.DeleteTask(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task Not Found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task Deleted"))
}
