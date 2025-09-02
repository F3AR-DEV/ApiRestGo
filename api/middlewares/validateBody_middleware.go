package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/F3AR-DEV/ApiRestGO/utils"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateBodyMiddleware(next http.Handler, dtoType interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Crear una nueva instancia del DTO
		payload := reflect.New(reflect.TypeOf(dtoType)).Interface()

		// Decodificar body en el struct proporcionado
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.WriteJSONResponse(w, http.StatusBadRequest, "Error", "Bad request", nil)
			return
		}

		// Validar struct usando tags `validate`
		if err := validate.Struct(payload); err != nil {
			utils.WriteJSONResponse(w, http.StatusBadRequest, "Error", "Bad request", nil)
			return
		}

		// Guardar payload validado en context
		ctx := context.WithValue(r.Context(), ValidatedBodyKey, payload)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
