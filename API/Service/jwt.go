package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Clave secreta para firmar y verificar tokens temporal
var claveSecreta = []byte("mi_clave_secreta")

func CreateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token.Claims = claims

	tokenString, err := token.SignedString(claveSecreta)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return claveSecreta, nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("token no válido")
	}

	return nil
}
