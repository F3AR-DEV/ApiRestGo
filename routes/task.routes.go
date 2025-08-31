package routes

import (
	"encoding/json"
	"net/http"

	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)

	w.Write([]byte("Get tasks"))
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("task No found"))
		return
	}

	json.NewEncoder(w).Encode(task)
	w.Write([]byte("Get task"))
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	// Debug: imprime la variable task en la consola del servidor
	jsonTask, _ := json.MarshalIndent(task, "", "  ")
	println(string(jsonTask))
	CreateTask := db.DB.Create(&task)
	err := CreateTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.User
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("task No found"))
		return
	}

	db.DB.Delete(&task)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete task"))
}
