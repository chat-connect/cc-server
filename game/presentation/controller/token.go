package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func CheckToken(baseToken string) (userKey string, name string, email string, err error) {
	token, err := jwt.Parse(baseToken[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		
		return []byte("secret"), nil
	})
	if err != nil {
		return userKey, name, email, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userKey, name, email, err
	}

	userKey = claims["user_key"].(string)
	name = claims["name"].(string)
	email = claims["email"].(string)

	return userKey, name, email, nil
}
