package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/F3AR-DEV/ApiRestGO/dto"
	"github.com/F3AR-DEV/ApiRestGO/middlewares"
	"github.com/F3AR-DEV/ApiRestGO/models"
	"github.com/F3AR-DEV/ApiRestGO/services"
	"github.com/F3AR-DEV/ApiRestGO/utils"
	"github.com/gorilla/mux"
)

// CreateItem crea un nuevo item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	// Recuperar payload validado desde el context
	req := r.Context().Value(middlewares.ValidatedBodyKey).(*dto.ItemRequest)

	item := models.Item{
		Item: req.Item,
	}

	if err := services.CreateItem(&item); err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, "service_error", "Failed to create item", nil)
		return
	}

	res := dto.ItemResponse{
		ID:   item.ID,
		Item: item.Item,
	}

	utils.WriteJSONResponse(w, http.StatusCreated, "controller_success", "Item created successfully", res)
}

// GetItems devuelve todos los items
func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := services.GetItems()
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, "service_error", "Failed to fetch items", nil)
		return
	}

	// Mapear a DTOs
	var responseItems []dto.ItemResponse
	for _, item := range items {
		responseItems = append(responseItems, dto.ItemResponse{
			ID:   item.ID,
			Item: item.Item,
		})
	}

	utils.WriteJSONResponse(w, http.StatusOK, "success", "Items retrieved successfully", responseItems)
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
