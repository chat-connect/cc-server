package controller

import (
	"fmt"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

func CheckToken(tokenString string) error {
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	if tokenString == "" {
		return fmt.Errorf("Invalid token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
		return []byte("secret"), nil
	})
	if err != nil {
		return fmt.Errorf("Invalid token")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return fmt.Errorf("Invalid token")
	}
}
