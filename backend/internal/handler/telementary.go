package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type TelemetryType struct {
	DB *db.Queries
}

// Issue represents a single issue in any category
type Issue struct {
	Type     string `json:"type"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
}

type TelemetryReport struct {
	AgentID     int64       `json:"agent_id"`
	AgentToken  string      `json:"agent_token"`
	Severity    string      `json:"severity"`
	RawPayload  []byte      `json:"raw_payload"`
	Status      string      `json:"status"`
	RiskLevel   pgtype.Text `json:"risk_level"`
	Message     pgtype.Text `json:"message"`
	Summary     pgtype.Text `json:"summary"`
	Performance []Issue     `json:"performance"`
	Network     []Issue     `json:"network"`
	Security    []Issue     `json:"security"`
	AlertType   string      `json:"alert_type"`
}

func (h *TelemetryType) ReceiveTelemetry(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var IncomingAgentInfo TelemetryReport
	err := json.NewDecoder(r.Body).Decode(&IncomingAgentInfo)

	utils.CheckError(w, err, http.StatusBadGateway, "invalid info")

	user, err := h.DB.GetAgentByAgentId(ctx, IncomingAgentInfo.AgentID)
	utils.CheckError(w, err, http.StatusBadGateway, "internal server error")

	// validate AgentID and AgentToken
	if IncomingAgentInfo.AgentID != user.AgentID && IncomingAgentInfo.AgentToken != user.AgentToken {
		return
	}

	perfJSON, err := json.Marshal(IncomingAgentInfo.Performance)
	utils.CheckError(w, err, http.StatusInternalServerError, "internal server error")
	netJSON, err := json.Marshal(IncomingAgentInfo.Network)
	utils.CheckError(w, err, http.StatusInternalServerError, "internal server error")
	secJSON, err := json.Marshal(IncomingAgentInfo.Security)
	utils.CheckError(w, err, http.StatusInternalServerError, "internal server error")

	AlertToCommit := db.CreateAlertParams{
		AgentID:     user.UserID,
		AlertType:   IncomingAgentInfo.AlertType,
		Severity:    IncomingAgentInfo.Severity,
		Message:     IncomingAgentInfo.Message,
		RawPayload:  IncomingAgentInfo.RawPayload,
		Status:      IncomingAgentInfo.Status,
		RiskLevel:   IncomingAgentInfo.RiskLevel,
		Summary:     IncomingAgentInfo.Summary,
		Performance: perfJSON,
		Network:     netJSON,
		Security:    secJSON,
	}
	fmt.Println(AlertToCommit)

	// store in database
	err = h.DB.CreateAlert(ctx, AlertToCommit)
	utils.CheckError(w, err, http.StatusInternalServerError, "internal server error")
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "alert created Successfully",
	})
}
