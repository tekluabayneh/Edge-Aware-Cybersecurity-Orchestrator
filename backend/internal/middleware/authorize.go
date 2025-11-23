package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func Authorize(db *db.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			isValid, token := utils.VerifyToken(r)
			if !isValid {
				utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
					"message": "invalid token",
				})
				return
			}
			claims := token.Claims.(jwt.MapClaims)
			expiryTime := time.Unix(int64(claims["exp"].(float64)), 0)
			email := claims["user_email"]
			emailStr, ok := email.(string)

			if !ok {
				utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
					"message": "invalid email in token",
				})
				return
			}

			if time.Now().After(expiryTime) {
				utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
					"message": "token is expired",
				})
				return
			}
			user, err := db.GetUserByEmail(ctx, emailStr)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
					"message": "internal server error",
				})
				return
			}

			if err != nil && user.Email == "" {
				utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
					"message": "UnAuthorized user",
				})
				return
			}
			ctxValue := context.WithValue(ctx, "email", email)
			next.ServeHTTP(w, r.WithContext(ctxValue))

		})
	}
}
