package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId string, username string) (string, error) {
	jwtSecret := LoadEnvVariables().AccessTokenSecret

	secretKey := []byte(jwtSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["username"] = username

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
