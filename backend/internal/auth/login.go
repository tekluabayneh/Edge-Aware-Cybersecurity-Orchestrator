package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/middleware"
	"github.com/edge-aware-cyberSecurity/internal/utils"
)

type AuthLoginHandlerType struct {
	DB *db.Queries
}

type LoginStructure struct {
	Name  string
	Email string
}

func (h *AuthLoginHandlerType) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	ctx := r.Context()

	formData, ok := r.Context().Value(middleware.LoginFromDataKey).(middleware.FormType)
	if !ok {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	user, err := h.DB.GetUserByEmail(ctx, formData.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	if err != nil && user.Email == "" {
		// here user is does not have account
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "user does not exists",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "password does not match please try again",
		})
		return
	}

	fmt.Println(err)
	token, err := utils.GenerateToken(formData.Email)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "error generating token",
			"detail":  err.Error(),
		})
		return
	}

	// send success message with token
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "user login successfully",
		"token":   token,
	})
}
