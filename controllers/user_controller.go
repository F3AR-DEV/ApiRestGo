package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/F3AR-DEV/ApiRestGO/services"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := services.GetUserByID(params["id"])
	if err != nil || user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	json.NewEncoder(w).Encode(user)
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	err := services.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := services.DeleteUser(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Deleted"))
}
