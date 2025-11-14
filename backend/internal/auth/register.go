package handler

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/middleware"
	"github.com/edge-aware-cyberSecurity/internal/utils"
)

type AuthRegisterHandlerType struct {
	DB *db.Queries
}

type RegisterStructure struct {
	Name     string
	Email    string
	Password string
}

func (h *AuthRegisterHandlerType) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	// get user data form context
	formData, ok := r.Context().Value(middleware.RegisterFormDataKey).(middleware.FormType)
	if !ok {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	// check if user already has account
	user, err := h.DB.GetUserByEmail(ctx, formData.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	if err == nil && user.Email != "" {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "user already exist",
		})
		return
	}

	HashedPassword, err := utils.HashPassword(formData.Password)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	userInfo := db.CreateUserParams{
		Name:     formData.Name,
		Email:    formData.Email,
		Password: HashedPassword,
	}

	// register user if does not exists
	err = h.DB.CreateUser(ctx, userInfo)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "user registered successfully",
	})
}
