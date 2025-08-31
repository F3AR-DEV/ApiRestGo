package routes

import (
	"encoding/json"

	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/gorilla/mux"

	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

	w.Write([]byte("Get Users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User No found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
	w.Write([]byte("Get User"))
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser := db.DB.Create(&user)
	err := CreateUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)

	w.Write([]byte("Post"))
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User No found"))
		return
	}

	db.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete User"))
}
