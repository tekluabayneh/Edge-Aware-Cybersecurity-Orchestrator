package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateCommdnType struct {
	DB *db.Queries
}

type incomingPayload struct {
	Email string `json:"email"`
}

type AckResponseType struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
	IsAck   bool   `json:"is_ack"`
}

type commandPaylod struct {
	UserId      int64              `json:"user_id"`
	Email       string             `json:"email"`
	AgentId     string             `json:"agent_id"`
	CommandType string             `json:"command_type"`
	Payload     json.RawMessage    `json:"payload"`
	Status      string             `json:"status"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type notification struct {
	UserId    int64              `json:"user_id"`
	Title     string             `json:"title"`
	Message   string             `json:"message"`
	IsRead    bool               `json:"is_read"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

// Create a new command for a device (from dashboard)
func (h *CreateCommdnType) CreateCommandHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var Email incomingPayload
	json.NewDecoder(r.Body).Decode(&Email)

	user, err := h.DB.GetUserByEmail(ctx, Email.Email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "user not found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusNotFound, "user not found") {
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

	var jsonPaylod map[string]any
	err = json.Unmarshal(IncomPayloadtype.Payload, &jsonPaylod)

	if utils.CheckError(w, err, http.StatusBadRequest, "paylod is not correct") {
		return
	}

	params := db.CreateCommandParams{
		UserID:      int64(user.ID),
		AgentID:     IncomPayloadtype.AgentId,
		CommandType: IncomPayloadtype.CommandType,
		Payload:     IncomPayloadtype.Payload,
		Status:      IncomPayloadtype.Status,
		UpdatedAt:   CreateAtTime,
	}

	err = h.DB.CreateCommand(ctx, params)
	if utils.CheckError(w, err, http.StatusNotFound, "error committing command") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "commadn created successfully",
	})

}

// Fetch pending commands for a given agent
func (h *CreateCommdnType) FetchPendingCommandsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var Email incomingPayload
	json.NewDecoder(r.Body).Decode(&Email)

	user, err := h.DB.GetUserByEmail(ctx, Email.Email)
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
	if utils.CheckError(w, err, http.StatusNotFound, "agent not found") {
		return
	}

	command, err := h.DB.FetchPendingCommndByAgentId(ctx, agent.MachineID)
	if utils.CheckError(w, err, http.StatusNotFound, "command not found") {
		return
	}

	var jsonpaylod map[string]any
	for _, cmd := range command {
		err = json.Unmarshal(cmd.Payload, &jsonpaylod)
		if utils.CheckError(w, err, http.StatusNotFound, "invalid paylod form") {
			return
		}
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "command Fetched successfully",
		"command": command,
		"payload": jsonpaylod,
	})

}

// Acknowledge command execution from agent
func (h *CreateCommdnType) AcknowledgeCommandExecutionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var Email incomingPayload
	json.NewDecoder(r.Body).Decode(&Email)

	user, err := h.DB.GetUserByEmail(ctx, Email.Email)
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

	status, err := h.DB.UpdateCommandStatusByAgentId(ctx, agent.MachineID)
	if utils.CheckError(w, err, http.StatusNotFound, "command not found") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "commadn updated ruccessfully",
		"status":  status,
	})

}

func (h *CreateCommdnType) AcknowledgeCommandExecutionResponseHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var AckREsponse AckResponseType
	json.NewDecoder(r.Body).Decode(&AckREsponse)
	user, err := h.DB.GetUserByEmail(ctx, AckREsponse.Email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "user not found",
		})
		return
	}

	notificationPayload := db.CreateNotificationParams{
		UserID:  int64(user.ID),
		Message: AckREsponse.Message,
		Title:   "check for the command acknowledge",
		IsRead:  false,
	}

	err = h.DB.CreateNotification(ctx, notificationPayload)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "user not found",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "commadn updated ruccessfully",
	})
}
