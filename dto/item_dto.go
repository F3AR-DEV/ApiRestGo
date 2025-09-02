package dto

// ItemRequest se usa para crear o actualizar un Item
type ItemRequest struct {
	Item string `json:"item" validate:"required,min=3,max=128"`
}

// ItemResponse se usa para devolver datos al cliente
type ItemResponse struct {
	ID   uint   `json:"id"`
	Item string `json:"item"`
}
