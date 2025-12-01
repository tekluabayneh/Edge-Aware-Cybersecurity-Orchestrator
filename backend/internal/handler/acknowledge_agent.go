package handler

import (
	"encoding/json"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type Agent struct {
	UserID       int64              `json:"user_id"`
	AgentToken   string             `json:"agent_token"`
	AgentId      int64              `json:"agent_id"`
	MachineID    string             `json:"machine_id"`
	AgentVersion pgtype.Text        `json:"agent_version"`
	Os           pgtype.Text        `json:"os"`
	Status       pgtype.Text        `json:"status"`
	LastSeen     pgtype.Timestamptz `json:"last_seen"`
}

func (h *DevicePairingType) AcknowledgePairing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var AgentValue Agent
	err := json.NewDecoder(r.Body).Decode(&AgentValue)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"message": "invalid request",
			"ack":     false,
		})
		return
	}

	if !utils.RequiredFieldsSet(utils.Agent(AgentValue)) {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"message": "missing required info",
			"ack":     false,
		})
		return
	}

	AgentInfo := db.CreateAgentParams{
		UserID:       AgentValue.UserID,
		AgentToken:   AgentValue.AgentToken,
		AgentID:      AgentValue.AgentId,
		MachineID:    AgentValue.MachineID,
		AgentVersion: AgentValue.AgentVersion,
		Os:           AgentValue.Os,
		Status:       AgentValue.Status,
		LastSeen:     AgentValue.LastSeen,
	}

	err = h.DB.CreateAgent(ctx, AgentInfo)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"ack":     false,
			"dl":      err.Error(),
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "agent info created successfully",
		"ack":     true,
	})

}
