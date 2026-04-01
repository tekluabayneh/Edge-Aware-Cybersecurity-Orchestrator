package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"user_email": email,
		"exp":        time.Now().Add(time.Hour * 14).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateDeviceParingToken(length int) string {
	token := uuid.New().String()
	return token
}
