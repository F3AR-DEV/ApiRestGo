package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ParseAndValidate decodifica y valida el body contra un struct destino
func ParseAndValidate(r *http.Request, dst interface{}) error {
	// Decodificar JSON
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return err
	}

	// Validar estructura
	if err := validate.Struct(dst); err != nil {
		return err
	}

	return nil
}
