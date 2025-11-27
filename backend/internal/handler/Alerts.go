package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/go-chi/chi/v5"
)

type AlertType struct {
	DB *db.Queries
}

/*
	ALERT ENDPOINTS:

	GET    /api/alerts              -> Get all alerts
	GET    /api/alerts/:agent_id    -> Get alerts by agent ID
	PATCH  /api/alerts/:id/read     -> Mark single alert as read
	PATCH  /api/alerts/read-all     -> Mark all alerts as read
	DELETE /api/alerts/:id          -> Delete an alert
	GET    /api/alerts/stats        -> Alert statistics
*/

// GET /api/alerts
func (h *AlertType) Alerts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	fmt.Println(user)
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}
	alerts, err := h.DB.GetAllAlert(ctx, int64(user.ID))
	if utils.CheckError(w, err, http.StatusBadRequest, "agent not found") {
		return
	}

	var rawJson []map[string]interface{}
	err = json.Unmarshal(alerts.RawPayload, &rawJson)
	if utils.CheckError(w, err, http.StatusBadRequest, "agent not found") {
		return
	}

	var SecurityJson []map[string]interface{}
	err = json.Unmarshal(alerts.Security, &SecurityJson)
	if utils.CheckError(w, err, http.StatusBadRequest, "agent not found") {
		return
	}

	var PerformanceJson []map[string]interface{}
	err = json.Unmarshal(alerts.Performance, &PerformanceJson)
	if utils.CheckError(w, err, http.StatusBadRequest, "agent not found") {
		return
	}

	var networkJson []map[string]interface{}
	err = json.Unmarshal(alerts.Network, &networkJson)
	if utils.CheckError(w, err, http.StatusBadRequest, "agent not found") {
		return
	}

	/* AlertToSend = map[string]interface{}{
		"network":     networkJson,
		"performance": PerformanceJson,
		"security":    SecurityJson,
		"raw":         rawJson,
		"alertType":   alerts.AlertType,
		"agent_id":    alerts.agent_id,
		"message":     alerts.Message,
		"summery":     alerts.Security,
		"agent_token": alerts.agent_token,
		"risk_level":  alerts.RiskLevel,
		"status":      alerts.Status,
	}
	fmt.Println(AlertToSend) */

	// Encode the actual structured data
	// json.NewEncoder(w).Encode(AlertToSend)
	//get agent_id or email if not the agent_id
	// fetch all alerted related to user

}

// GET /api/alerts/:agent_id
func (h *AlertType) GetAlertByAgentId(w http.ResponseWriter, r *http.Request) {
	agentID := chi.URLParam(r, "agent_id")
	fmt.Println("→ Get alerts for agent:", agentID)
	// TODO: fetch alerts by agent ID from DB
	w.Write([]byte("get alerts by agent_id"))
}

// PATCH /api/alerts/:id/read
func (h *AlertType) UpdateAlertsById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	alertID, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		http.Error(w, "invalid alert id", http.StatusBadRequest)
		return
	}

	fmt.Println("→ Mark alert as read:", alertID)

	// TODO: mark alert as read in DB
	w.Write([]byte("updated alert by id"))
}

// PATCH /api/alerts/read-all
func (h *AlertType) UpdateAllAlerts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("→ Mark all alerts as read")

	// TODO: mark all alerts as read in DB
	w.Write([]byte("updated all alerts"))
}

// DELETE /api/alerts/:id
func (h *AlertType) DeleteAlertById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	alertID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid alert id", http.StatusBadRequest)
		return
	}

	fmt.Println("→ Delete alert:", alertID)

	// TODO: delete alert from DB
	w.Write([]byte("deleted alert by id"))
}

// GET /api/alerts/stats
func (h *AlertType) GetAllAlertStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("→ Get alert statistics")

	// Example response structure
	stats := map[string]int{
		"total":       100, // TODO: replace with real DB query
		"unread":      25,
		"read":        75,
		"high_risk":   10,
		"medium_risk": 50,
		"low_risk":    40,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
