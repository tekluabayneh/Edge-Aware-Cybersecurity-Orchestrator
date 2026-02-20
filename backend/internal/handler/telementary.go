package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type Issue struct {
	Type     string `json:"type" validate:"required"`
	Message  string `json:"message" validate:"required"`
	Severity string `json:"severity" validate:"required"`
}

type TelemetryReport struct {
	AgentID     string             `json:"agent_id" validate:"required"`
	AgentToken  string             `json:"agent_token" validate:"required"`
	Severity    string             `json:"severity" validate:"required"`
	RawPayload  []Issue            `json:"raw_payload" validate:"required,dive,required"`
	Status      string             `json:"status" validate:"required"`
	Message     pgtype.Text        `json:"message" validate:"required"`
	RiskLevel   pgtype.Text        `json:"risk_level" validate:"required"`
	Summary     pgtype.Text        `json:"summary" validate:"required"`
	Performance []Issue            `json:"performance" validate:"required,dive,required"`
	Network     []Issue            `json:"network" validate:"required,dive,required"`
	Security    []Issue            `json:"security" validate:"required,dive,required"`
	AlertType   string             `json:"alert_type" validate:"required"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}
type TelemetryType struct {
	DB *db.Queries
}

func (h *TelemetryType) ReceiveTelemetry(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var IncomingAgentInfo TelemetryReport
	err := json.NewDecoder(r.Body).Decode(&IncomingAgentInfo)
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid telemetry payload") {
		return
	}

	agent, err := h.DB.GetAgentByAgentToken(ctx, IncomingAgentInfo.AgentToken)
	if utils.CheckError(w, err, http.StatusBadGateway, "agent not found") {
		return
	}

	if IncomingAgentInfo.AgentToken != agent.AgentToken {
		utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
			"message": "invalid agent credentials",
		})
		return
	}

	perfJSON, err := json.Marshal(IncomingAgentInfo.Performance)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed performance encoding") {
		return
	}

	netJSON, err := json.Marshal(IncomingAgentInfo.Network)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed network encoding") {
		return
	}

	secJSON, err := json.Marshal(IncomingAgentInfo.Security)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed security encoding") {
		return
	}

	rawJSON, err := json.Marshal(IncomingAgentInfo.RawPayload)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed raw_payload encoding") {
		return
	}

	CreateAtTime := pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}

	AlertToCommit := db.CreateAlertParams{
		AgentID:     agent.MachineID,
		AgentToken:  IncomingAgentInfo.AgentToken,
		AlertType:   IncomingAgentInfo.AlertType,
		Severity:    IncomingAgentInfo.Severity,
		Message:     IncomingAgentInfo.Message,
		RawPayload:  rawJSON,
		Status:      IncomingAgentInfo.Status,
		RiskLevel:   IncomingAgentInfo.RiskLevel,
		Summary:     IncomingAgentInfo.Summary,
		Performance: perfJSON,
		Network:     netJSON,
		Security:    secJSON,
		CreatedAt:   CreateAtTime,
	}

	err = h.DB.CreateAlert(ctx, AlertToCommit)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed creating alert") {
		fmt.Println(err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "telemetry created successfully",
	})

}
