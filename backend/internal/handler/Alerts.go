package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	notify "github.com/edge-aware-cyberSecurity/internal/handler/notification"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type AlertType struct {
	DB *db.Queries
}

type alertStatusType struct {
	Status string `json:"status"`
}

type NewAlertRequest struct {
	AgentID     string                 `json:"agent_id"`
	Email       string                 `json:"email"`
	AgentToken  string                 `json:"agent_token"`
	AlertType   string                 `json:"alert_type"`
	Severity    string                 `json:"severity"`
	Message     pgtype.Text            `json:"message"`
	RawPayload  map[string]interface{} `json:"raw_payload"`
	Status      string                 `json:"status"`
	RiskLevel   pgtype.Text            `json:"risk_level"`
	Summary     pgtype.Text            `json:"summary"`
	Performance map[string]interface{} `json:"performance"`
	Network     map[string]interface{} `json:"network"`
	Security    map[string]interface{} `json:"security"`
}

// GET /api/alerts
func (h *AlertType) Alerts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}
	agent, err := h.DB.GetAgentByUserId(ctx, int64(user.ID))
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "agent not found so No Alert",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errors") {
		return
	}

	// Get all alerts for the user
	alerts, err := h.DB.GetAllAlert(ctx, string(agent.AgentID))
	if len(alerts) < 1 {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{
			"message": "alert not found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}

	var allAlerts []map[string]any
	for _, alert := range alerts {
		fields := []string{"Network", "Performance", "Security", "RawPayload"}
		unmarshaledData := make(map[string]any)

		val := reflect.Indirect(reflect.ValueOf(alert))

		for _, field := range fields {
			f := val.FieldByName(field)
			if !f.IsValid() {
				continue
			}

			fieldBytes, ok := f.Interface().([]byte)
			if !ok || len(fieldBytes) == 0 {
				continue
			}

			var tmp any
			if err := json.Unmarshal(fieldBytes, &tmp); err != nil {
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
			"CreatedAt":   alert.CreatedAt,
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
		"AgentID":     alert.AgentID,
		"AgentToken":  alert.AgentToken,
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

func (h *AlertType) CreateAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	email := r.Context().Value("email").(string)
	var Alerts NewAlertRequest
	err := json.NewDecoder(r.Body).Decode(&Alerts)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error crashed whie decoding body") {
		return
	}

	user, err := h.DB.GetUserByEmail(ctx, Alerts.Email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		if utils.CheckError(w, err, http.StatusNotFound, "user not found for the agent") {
			return
		}
	}

	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error!") {
		return
	}

	agent, err := h.DB.GetAgentByUserId(ctx, int64(user.ID))
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{
			"message": "agent not found!",
		})
		return
	}

	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error!!") {
		return
	}

	if agent.AgentID != Alerts.AgentID {
		if utils.CheckError(w, err, http.StatusInternalServerError, "incoming agent id didn't match") {
			return
		}
	}
	rawJsonpaylod, err := json.Marshal(Alerts.RawPayload)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error, faild to marshal alert rawpaylod") {
		return
	}

	jsonPerformance, err := json.Marshal(Alerts.Performance)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error, faild to marshal alert performance") {
		return
	}

	jsonNetwork, err := json.Marshal(Alerts.Network)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error, faild to marshal alert netwrok") {
		return
	}

	jsonSecurity, err := json.Marshal(Alerts.Security)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error, faild to marshal alert security") {
		return
	}

	params := db.CreateAlertParams{
		AgentID:     Alerts.AgentID,
		AgentToken:  Alerts.AgentToken,
		Message:     Alerts.Message,
		AlertType:   Alerts.AlertType,
		Severity:    Alerts.Severity,
		RawPayload:  rawJsonpaylod,
		Status:      Alerts.Status,
		RiskLevel:   Alerts.RiskLevel,
		Summary:     Alerts.Summary,
		Performance: jsonPerformance,
		Network:     jsonNetwork,
		Security:    jsonSecurity,
	}

	err = h.DB.CreateAlert(ctx, params)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}

	// and also chech if user allowd
	if strings.ToLower(params.Severity) == "critical" && user.Notification.Bool {
		notificationParams := db.CreateNotificationParams{
			UserID:  int64(user.ID),
			Title:   Alerts.Summary.String,
			Message: Alerts.Message.String,
			IsRead:  false,
		}

		err = h.DB.CreateNotification(ctx, notificationParams)
		if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
			return
		}

		smtRes, res, err := notify.Notify(email, email, Alerts.Summary.String)
		if err != nil {
			fmt.Println(smtRes)
			fmt.Println(res)
			fmt.Println("err sending alert email")
		}
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "alert create successfully",
	})
}
