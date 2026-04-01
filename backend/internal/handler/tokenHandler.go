package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

// when user want to generate connection token we first must handler edge cases
// first make sure this person exist
// second check if the person has more thatn 10 agent connected
// thrid  check if the person has ganrated token if so and if the person requteed for the new one remove the old oe and update it with the new one and send it to client

func (h *DevicePairingType) GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	UserEmail := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, UserEmail)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{
			"message": "user not found",
		})
		return
	}

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	t := time.Now().Add(6 * time.Hour)
	expires := pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}

	token := utils.GenerateDeviceParingToken(16)
	DeviceCount, err := h.DB.GetUserDeviceCount(ctx, int64(user.ID))
	// if user does not have any agent generate teh code and return early
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		paringDeviceData := db.CreateParingTokenParams{
			Token:     token,
			UserID:    user.ID,
			UserEmail: user.Email,
			ExpiresAt: expires,
		}

		err = h.DB.CreateParingToken(ctx, paringDeviceData)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		utils.WriteJSON(w, http.StatusOK, map[string]string{
			"message": "paring token created",
			"token":   token,
		})

		return
	}

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	if DeviceCount > 10 {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Device paring is limited you have gone your free limit",
		})
		return
	}

	// get token
	AllUserToken, err := h.DB.GetAllTokenRelatedToUserByEmail(ctx, user.Email)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})

		return
	}

	// if token exists → delete it
	if err == nil {
		// optional: check expiration safely
		if AllUserToken.ExpiresAt.Valid && time.Now().After(AllUserToken.ExpiresAt.Time) {
			fmt.Println("Token expired")
		}

		err = h.DB.DeleteUsedToken(ctx, UserEmail)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}
	}

	// ALWAYS create new token
	paringDeviceData := db.CreateParingTokenParams{
		Token:     token,
		UserID:    user.ID,
		UserEmail: user.Email,
		ExpiresAt: expires,
	}

	err = h.DB.CreateParingToken(ctx, paringDeviceData)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "paring token created",
		"token":   token,
	})
}
