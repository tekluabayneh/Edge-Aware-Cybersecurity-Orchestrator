package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
)

type DevicePairingType struct {
	DB *db.Queries
}

func (h *DevicePairingType) DevicePairing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	UserEmail := r.Context().Value("email").(string)
	incomingToken := r.URL.Query().Get("token")

	if incomingToken == "" {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "token not found",
		})
		return
	}

	userToken, err := h.DB.GetAllParingTokenByEmail(ctx, UserEmail)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "internal server error",
		})
		return
	}

	if err == nil && len(userToken.Token) < 1 {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "token not found",
		})
		return
	}

	if time.Now().After(userToken.ExpiresAt.Time) {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "token is expired",
		})
		return
	}

	if userToken.Token != incomingToken {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid token",
		})
		return
	}

	err = h.DB.DeleteUsedToken(ctx, UserEmail)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "internal server error",
		})
		return
	}

	agent_id := utils.GenerateDeviceParingToken(6)
	agent_token := utils.GenerateDeviceParingToken(16)

	// send agent_id and agent_token with success message and agent will send agent info and stor them
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message":     "token validation success",
		"agent_id":    agent_id,
		"agent_token": agent_token,
	})
}
