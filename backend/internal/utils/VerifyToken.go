package utils

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(r *http.Request) (bool, *jwt.Token) {
	incomingToken := r.Header.Get("Authorization")
	splitToken := strings.TrimPrefix(incomingToken, "Bearer ")
	token, err := jwt.Parse(splitToken, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err == nil && token.Valid {
		return true, token
	} else {
		return false, nil
	}
}
