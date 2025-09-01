package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/F3AR-DEV/ApiRestGO/services"
	"github.com/gorilla/mux"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateItem(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := services.GetItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func GetItemByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	item, err := services.GetItemByID(uint(id))
	if err != nil {
		http.Error(w, "Item no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	var newData models.Item
	if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := services.UpdateItem(uint(id), &newData)
	if err != nil {
		http.Error(w, "Error al actualizar item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	if err := services.DeleteItem(uint(id)); err != nil {
		http.Error(w, "Error al eliminar item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
