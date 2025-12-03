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
)

type AlertType struct {
	DB *db.Queries
}
type alertStatusType struct {
	Status string `json:"status"`
}

// GET /api/alerts
func (h *AlertType) Alerts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}
	// Get all alerts for the user
	alerts, err := h.DB.GetAllAlert(ctx, string(user.ID))
	if len(alerts) < 1 {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{
			"message": "alert not found",
		})
		return
	}
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

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "alerts fetched successfully",
		"alert":   allAlerts,
	})
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
		AgentID: string(user.ID),
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
	ctx := r.Context()
	id := r.URL.Query().Get("id")
	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid agent id") {
		return
	}
	if user.Email == "" {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "agent id not found",
		})
		return
	}
	alertID, err := strconv.ParseInt(id, 10, 64)
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid alert id") {
		return
	}
	var statusValue alertStatusType
	err = json.NewDecoder(r.Body).Decode(&statusValue)
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid alert status") {
		return
	}

	params := db.UpdateSingleAlertStatusParams{
		ID:      alertID,
		AgentID: string(user.ID),
		Status:  statusValue.Status,
	}
	_, err = h.DB.UpdateSingleAlertStatus(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "alert not found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid alert id") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "updated alert by id",
	})
}

// PATCH /api/alerts/read-all
func (h *AlertType) UpdateAllAlerts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid agent id") {
		return
	}
	if user.Email == "" {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "agent id not found",
		})
		return
	}
	var statusValue alertStatusType
	err = json.NewDecoder(r.Body).Decode(&statusValue)
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid alert status") {
		return
	}
	params := db.UpdateAllAlertStatusByAgentIdParams{
		AgentID: string(user.ID),
		Status:  statusValue.Status,
	}
	_, err = h.DB.UpdateAllAlertStatusByAgentId(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "alert not found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid alert id") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "updated alert by id",
	})
}

// GET /api/alerts/stats
func (h *AlertType) GetAllAlertStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}

	AlertSTatus, err := h.DB.GetAllAlertStatus(ctx, string(user.ID))
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}

	if len(AlertSTatus) < 1 {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "no alert found",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "alert status fetched successfully",
		"status":  AlertSTatus,
	})

}

// DELETE /api/alerts/:id
func (h *AlertType) DeleteAlertById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	idQuery := r.URL.Query().Get("id")
	user, err := h.DB.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}
	alertID, err := strconv.ParseInt(idQuery, 10, 64)
	if utils.CheckError(w, err, http.StatusBadRequest, "invalid alert id") {
		return
	}
	params := db.DeleteAlertByAGentIdParams{
		ID:      alertID,
		AgentID: string(user.ID),
	}
	DeltedAlertId, err := h.DB.DeleteAlertByAGentId(ctx, params)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": fmt.Sprintf("message not found with the id of %d", DeltedAlertId),
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "deleted alert by id",
	})

}
