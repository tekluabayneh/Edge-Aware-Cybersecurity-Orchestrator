package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
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
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}

	// Get all alerts for the user
	alerts, err := h.DB.GetAllAlert(ctx, int64(user.ID))
	if utils.CheckError(w, err, http.StatusBadRequest, "agent not found") {
		return
	}

	var allAlerts []map[string]any
	for _, alert := range alerts {
		fields := []string{"Network", "Performance", "Security", "RawPayload"}
		unmarshaledData := make(map[string][]map[string]any)

		val := reflect.ValueOf(alert)
		for _, field := range fields {
			f := val.FieldByName(field)
			if !f.IsValid() {
				continue
			}

			fieldBytes, ok := f.Interface().([]byte)
			if !ok {
				continue
			}

			var tmp []map[string]any
			err := json.Unmarshal(fieldBytes, &tmp)
			if utils.CheckError(w, err, http.StatusBadRequest, "agent not found") {
				return
			}

			unmarshaledData[field] = tmp
		}

		alertMap := map[string]any{
			"network":     unmarshaledData["Network"],
			"performance": unmarshaledData["Performance"],
			"security":    unmarshaledData["Security"],
			"raw":         unmarshaledData["RawPayload"],
			"alertType":   alert.AlertType,
			"agent_id":    alert.AgentID,
			"message":     alert.Message,
			"summery":     alert.Summary,
			"agent_token": alert.AgentToken,
			"risk_level":  alert.RiskLevel,
			"status":      alert.Status,
		}
		allAlerts = append(allAlerts, alertMap)
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "alerts fetched successfully",
	})
	json.NewEncoder(w).Encode(allAlerts)
}

// GET /api/alerts/:agent_id
func (h *AlertType) GetAlertByAgentId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.URL.Query().Get("id")
	email := r.Context().Value("email").(string)
	if userId == "" {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "params id is missing",
		})
		return
	}
	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if utils.CheckError(w, err, http.StatusBadRequest, "error parsing id") {
		return
	}

	params := db.GetAlertByAgentIdParams{
		ID:      id,
		AgentID: int64(user.ID),
	}
	alert, err := h.DB.GetAlertByAgentId(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{
			"message": "no alert found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "agent or id  not found") {
		return
	}

	networkJson := []map[string]any{}
	err = json.Unmarshal(alert.Network, &networkJson)
	if utils.CheckError(w, err, http.StatusBadRequest, "error json") {
		return
	}
	securityJson := []map[string]any{}
	err = json.Unmarshal(alert.Network, &securityJson)
	if utils.CheckError(w, err, http.StatusBadRequest, "error json") {
		return
	}
	performanceJson := []map[string]any{}
	err = json.Unmarshal(alert.Network, &performanceJson)
	if utils.CheckError(w, err, http.StatusBadRequest, "error json") {
		return
	}
	alertToSend := map[string]any{
		"ID":          alert.ID,
		"AgentID":     1,
		"AgentToken":  "abc123xyz",
		"AlertType":   alert.AlertType,
		"Severity":    alert.Severity,
		"Message":     alert.Message,
		"RawPayload":  alert.RawPayload,
		"Status":      alert.Status,
		"RiskLevel":   alert.RiskLevel,
		"Summary":     alert.Summary,
		"Performance": performanceJson,
		"Network":     networkJson,
		"Security":    securityJson,
		"CreatedAt":   alert.CreatedAt,
	}
	json.NewEncoder(w).Encode(alertToSend)
}

// PATCH /api/alerts/:id/read
func (h *AlertType) UpdateAlertsById(w http.ResponseWriter, r *http.Request) {
	idQuiery := r.URL.Query().Get("id")
	alertID, err := strconv.ParseInt(idQuiery, 10, 64)

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
