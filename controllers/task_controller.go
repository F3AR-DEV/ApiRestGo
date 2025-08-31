package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/F3AR-DEV/ApiRestGO/services"
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
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	// Debug: imprime la task en consola
	jsonTask, _ := json.MarshalIndent(task, "", "  ")
	println(string(jsonTask))

	err := services.CreateTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
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
