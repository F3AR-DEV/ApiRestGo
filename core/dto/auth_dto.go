package dto

// RegisterRequest representa el payload para registrar un usuario
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginRequest representa el payload para login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserResponse representa la respuesta de un usuario
type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
