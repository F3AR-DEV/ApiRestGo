package controllers

import (
	"fmt"
	"net/http"

	"github.com/F3AR-DEV/ApiRestGO/api/middlewares"
	"github.com/F3AR-DEV/ApiRestGO/core/dto"
	"github.com/F3AR-DEV/ApiRestGO/core/services"
	"github.com/F3AR-DEV/ApiRestGO/data/models"
	"github.com/F3AR-DEV/ApiRestGO/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Recuperar payload validado desde el context
	fmt.Println("Hola mundo")
	req := r.Context().Value(middlewares.ValidatedBodyKey).(*dto.RegisterRequest)

	// Hashear contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, "hash_error", "Failed to hash password", nil)
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Guardar usuario
	if err := services.CreateUser(&user); err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, "service_error", "Failed to create user", nil)
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, "success", "User registered successfully", dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}

// LoginHandler maneja el login y devuelve JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Recuperar payload validado desde el context
	req := r.Context().Value(middlewares.ValidatedBodyKey).(*dto.LoginRequest)

	user, err := services.GetUserByEmail(req.Email)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusUnauthorized, "auth_error", "Invalid email or password", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.WriteJSONResponse(w, http.StatusUnauthorized, "auth_error", "Invalid email or password", nil)
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, "jwt_error", "Failed to generate token", nil)
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, "controller_success", "Login successful", map[string]string{"token": token})
}
