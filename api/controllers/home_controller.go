package controllers

import (
	"net/http"

	"github.com/F3AR-DEV/ApiRestGO/business/services"
)

// HomeHandler maneja la petición al home
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	message := services.GetHomeMessage()
	w.Write([]byte(message))
}
