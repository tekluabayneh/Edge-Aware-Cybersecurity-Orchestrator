package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 1. Get user from context
	email, ok := r.Context().Value("email").(string)
	if !ok {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid user context",
		})
		return
	}

	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}

	// 2. Get current agent
	agent, err := h.DB.GetAgentByUserId(ctx, int64(user.ID))
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "no agent found",
			"alert":   []any{},
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}

	// 3. Try current agent alerts FIRST
	alertsToSend, err := h.DB.GetAllAlert(ctx, string(agent.AgentID))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "failed to fetch alerts",
		})
		return
	}
	fmt.Println(alertsToSend)
	// 4. If current agent has no alerts → check other agents
	if len(alertsToSend) == 0 {
		allAgents, err := h.DB.GetAllAgentByUserId(ctx, int64(user.ID))
		if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
			return
		}

		for _, v := range allAgents {
			// skip current agent
			// if v.AgentID == agent.AgentID {
			// 	continue
			// }

			alerts, err := h.DB.GetAllAlert(ctx, string(v.AgentID))

			fmt.Println("alerts", alerts)

			if errors.Is(err, sql.ErrNoRows) {
				continue
			}

			if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
				return
			}

			if len(alerts) > 0 {
				alertsToSend = alerts
				break
			}
		}
	}

	// 5. If still no alerts → return empty array
	if len(alertsToSend) == 0 {
		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "no alerts found",
			"alert":   []any{},
		})
		return
	}

	// 6. Transform alerts
	var allAlerts []map[string]any

	for _, alert := range alertsToSend {
		unmarshaledData := make(map[string]any)

		// Network
		if len(alert.Network) > 0 {
			var tmp any
			if err := json.Unmarshal(alert.Network, &tmp); err == nil {
				unmarshaledData["network"] = tmp
			}
		}

		// Performance
		if len(alert.Performance) > 0 {
			var tmp any
			if err := json.Unmarshal(alert.Performance, &tmp); err == nil {
				unmarshaledData["performance"] = tmp
			}
		}

		// Security
		if len(alert.Security) > 0 {
			var tmp any
			if err := json.Unmarshal(alert.Security, &tmp); err == nil {
				unmarshaledData["security"] = tmp
			}
		}

		// RawPayload
		if len(alert.RawPayload) > 0 {
			var tmp any
			if err := json.Unmarshal(alert.RawPayload, &tmp); err == nil {
				unmarshaledData["raw"] = tmp
			}
		}

		alertMap := map[string]any{
			"network":     unmarshaledData["network"],
			"performance": unmarshaledData["performance"],
			"security":    unmarshaledData["security"],
			"raw":         unmarshaledData["raw"],
			"alertType":   alert.AlertType,
			"agent_id":    alert.AgentID,
			"message":     alert.Message,
			"summary":     alert.Summary,
			"agent_token": alert.AgentToken,
			"risk_level":  alert.RiskLevel,
			"status":      alert.Status,
			"created_at":  alert.CreatedAt,
		}

		allAlerts = append(allAlerts, alertMap)
	}

	// 7. Final response
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "alerts fetched successfully",
		"alert":   allAlerts,
	})

}

// GET /api/alerts/:agent_id
func (h *AlertType) GetAlertByAgentId(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
	fmt.Println("hrere man", alertToSend)
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "alert fetched successfully",
		"alert":   alertToSend,
	})
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
