package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
)

type AgentType struct {
	DB *db.Queries
}

// GET all agent
func (h *AgentType) GetAllUserAgents(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}

	AllAgents, err := h.DB.GetAllAgentByUserId(ctx, int64(user.ID))
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "agent not found") {
		return
	}

	if len(AllAgents) < 1 {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "no alert found",
			"agents":  []string{},
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "alert status fetched successfully",
		"agents":  AllAgents,
	})

}
