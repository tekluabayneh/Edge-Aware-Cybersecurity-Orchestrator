package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *DevicePairingType) GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	UserEmail := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, UserEmail)

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	token := utils.GenerateDeviceParingToken(16)
	DeviceCount, err := h.DB.GetUserDeviceCount(ctx, int64(user.ID))
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

	// if user has token that is used or expired delete it
	AllUserToken, err := h.DB.GetAllTokenRelatedToUserByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
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

	if time.Now().After(AllUserToken.ExpiresAt.Time) {
		err = h.DB.DeleteUsedToken(ctx, UserEmail)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
				"message": "internal server error",
			})
			return
		}
	}

	if err == nil && len(AllUserToken.Token) > 0 {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "you already have token",
		})
		return
	}

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

	// send token to use paring dashboard
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "paring token created",
		"token":   token,
	})

}
