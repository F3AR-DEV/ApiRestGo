package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = []byte("mi_clave_secreta")

func GenerateJWT(email string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
