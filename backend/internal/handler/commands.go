package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateCommdnType struct {
	DB *db.Queries
}

type commandPaylod struct {
	UserId      int64              `json:"user_id"`
	AgentId     string             `json:"agent_id"`
	CommandType string             `json:"command_type"`
	Payload     []byte             `json:"payload"`
	Status      string             `json:"status"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

// Create a new command for a device (from dashboard)
func (h *CreateCommdnType) CreateCommandHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)

	user, err := h.DB.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "user not found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusNotFound, "user not found") {
		return
	}

	agent, err := h.DB.GetAgentByUserId(ctx, int64(user.ID))
	if utils.CheckError(w, err, http.StatusNotFound, "agent id not found") {
		return
	}

	var IncomPayloadtype commandPaylod
	err = json.NewDecoder(r.Body).Decode(&IncomPayloadtype)
	if utils.CheckError(w, err, http.StatusBadRequest, "paylod is not correct") {
		return
	}

	CreateAtTime := pgtype.Timestamptz{
		Time: time.Now(), Valid: true,
	}

	params := db.CreateCommandParams{
		UserID:      int64(user.ID),
		AgentID:     agent.AgentID,
		CommandType: IncomPayloadtype.CommandType,
		Payload:     IncomPayloadtype.Payload,
		Status:      IncomPayloadtype.Status,
		UpdatedAt:   CreateAtTime,
	}

	err = h.DB.CreateCommand(ctx, params)
	if utils.CheckError(w, err, http.StatusNotFound, "user not found") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "commadn created successfully",
	})

}

// Fetch pending commands for a given agent
func (h *CreateCommdnType) FetchPendingCommandsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)

	user, err := h.DB.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "user not found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusNotFound, "user not found") {
		return
	}

	agent, err := h.DB.GetAgentByUserId(ctx, int64(user.ID))
	if utils.CheckError(w, err, http.StatusNotFound, "user not found") {
		return
	}

	status, err := h.DB.FetchPendingCommndByAgentId(ctx, agent.AgentID)
	if utils.CheckError(w, err, http.StatusNotFound, "command not found") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "commadn Fetched successfully",
		"status":  status,
	})

}

// Acknowledge command execution from agent
func (h *CreateCommdnType) AcknowledgeCommandExecutionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)

	user, err := h.DB.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "user not found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusNotFound, "user not found") {
		return
	}

	agent, err := h.DB.GetAgentByUserId(ctx, int64(user.ID))
	if utils.CheckError(w, err, http.StatusNotFound, "user not found") {
		return
	}

	status, err := h.DB.UpdateCommandStatusByAgentId(ctx, agent.AgentID)
	if utils.CheckError(w, err, http.StatusNotFound, "user not found") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "commadn Fetched successfully",
		"status":  status,
	})

}
