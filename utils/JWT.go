package utils

import (
	"DzMart/initializers"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAndSignJwt(id uint) (string, error) {
	// Create New token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"ttl":    time.Now().Add(time.Hour * 24 * 100).Unix(),
	})
	Secret := initializers.EnvSecret()

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
