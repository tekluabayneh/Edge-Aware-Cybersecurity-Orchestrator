package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserType struct {
	DB *db.Queries
}

type UserResponse struct {
	ID                int32              `json:"id"`
	Name              string             `json:"name"`
	Email             string             `json:"email"`
	Photo             pgtype.Text        `json:"photo"`
	Phone             pgtype.Text        `json:"phone"`
	TwoFA             pgtype.Bool        `json:"two_fa"`
	Notification      pgtype.Bool        `json:"notification"`
	AlertNotification pgtype.Bool        `json:"alert_notification"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}

type UserUpdateType struct {
	Name               pgtype.Text        `json:"name"`
	Photo              pgtype.Text        `json:"photo"`
	Phone              pgtype.Text        `json:"phone"`
	Two_fa             pgtype.Bool        `json:"two_fa"`
	Notification       pgtype.Bool        `json:"notification"`
	Alert_notification pgtype.Bool        `json:"alert_notification"`
	Created_at         pgtype.Timestamptz `json:"created_at"`
}

func (h *UserType) GetUserProfileInfo(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := r.Context().Value("email").(string)

	user, err := h.DB.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, "user profile not found")
		return
	}

	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}

	resp := UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		Photo:             user.Photo,
		Phone:             user.Phone,
		TwoFA:             user.TwoFa,
		Notification:      user.Notification,
		AlertNotification: user.AlertNotification,
		CreatedAt:         user.CreatedAt,
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message":  "user profile fetched successfully",
		"UserInfo": resp,
	})
}

func (h *UserType) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := r.Context().Value("email").(string)

	// check if user exist
	oldProfile, err := h.DB.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "user profile not found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errror") {
		return
	}

	var NewInfo UserUpdateType
	json.NewDecoder(r.Body).Decode(&NewInfo)

	newINfo := db.UpdateUserProfileParams{
		Photo:             NewInfo.Photo,
		Phone:             NewInfo.Phone,
		Notification:      NewInfo.Notification,
		TwoFa:             NewInfo.Two_fa,
		Email:             oldProfile.Email,
		AlertNotification: NewInfo.Alert_notification,
	}

	err = h.DB.UpdateUserProfile(ctx, newINfo)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errror") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "user profile updated successfully",
	})

}
