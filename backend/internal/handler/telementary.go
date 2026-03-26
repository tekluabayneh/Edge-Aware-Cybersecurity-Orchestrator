package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

// --- NETWORK SECTION ---
type NetworkSection struct {
	Payload struct {
		ConnectionMonitoring struct {
			ActiveSockets      [][]ActiveSocket      `json:"ActiveSockets"`
			NetworkInterfaces  [][]NetworkInterface  `json:"NetworkInterfaces"`
			ConnectionPatterns [][]ConnectionPattern `json:"ConnectionPatterns"`
		} `json:"ConnectionMonitoring"`
	} `json:"payload"`
}

type ActiveSocket struct {
	FD             int          `json:"fd"`
	Family         int          `json:"family"`
	Type           int          `json:"type"`
	LocalAddr      AddrInfo     `json:"localaddr"`
	RemoteAddr     AddrInfo     `json:"remoteaddr"`
	Status         string       `json:"status"`
	UIDs           []int        `json:"uids"`
	PID            int          `json:"pid"`
	IsLocalhost    bool         `json:"is_localhost"`
	ConnectionType string       `json:"connection_type"`
	Suspicious     bool         `json:"suspicious"`
	RiskLevel      string       `json:"risk_level"`
	Rules          []SocketRule `json:"active_socker_rules"`
}

type AddrInfo struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type SocketRule struct {
	SocketListening string `json:"socket_listening"`
}

type NetworkInterface struct {
	Name           string          `json:"name"`
	Up             string          `json:"up"`
	Down           string          `json:"up"`
	IPAddresses    []IPAddr        `json:"ipAddresses"`
	InterfaceType  string          `json:"interface_type"`
	IsUpAndRunning bool            `json:"is_up_and_running"`
	IPCount        int             `json:"ip_count"`
	InterfaceRules []InterfaceRule `json:"networkInterface_rules"`
}

type IPAddr struct {
	Addr string `json:"addr"`
}

type InterfaceRule struct {
	Status   string `json:"interface_status"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
}

type ConnectionPattern struct {
	RemoteIP        string        `json:"remoteIp"`
	Frequency       int           `json:"frequency"`
	Volume          int           `json:"volume"`
	PatternType     string        `json:"pattern_type"`
	IsSuspiciousVol bool          `json:"is_suspicious_volume"`
	TrafficCategory string        `json:"traffic_category"`
	PotentialScan   bool          `json:"potential_scan"`
	PatternRules    []PatternRule `json:"connectionpatterns_rules"`
}

type PatternRule struct {
	Status  string `json:"connectionpattern_status"`
	Message string `json:"message"`
}

// --- SYSTEM, PROCESSES
type SystemSection struct {
	Payload map[string]interface{} `json:"payload"`
}

type ProcessesSection map[string]interface{}
type SecuritySection struct {
	Payload map[string]interface{} `json:"payload"`
}

type IntegritySection struct {
	Payload map[string]interface{} `json:"payload"`
}

type TelemetryReport struct {
	Email      string           `json:"email"`
	AgentId    string           `json:"agent_id"`
	AgentToken string           `json:"agent_token"`
	MachineID  string           `json:"machine_id"`
	System     SystemSection    `json:"system"`
	Security   SecuritySection  `json:"security"`
	Network    NetworkInterface `json:"network"`
	Processes  ProcessesSection `json:"processes"`
	Integrity  IntegritySection `json:"integrity"`
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

	securityJson, err := json.Marshal(IncomingAgentInfo.Security)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed security encoding") {
		return
	}

	processJson, err := json.Marshal(IncomingAgentInfo.Processes)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed processes encoding") {
		return
	}

	integrityJson, err := json.Marshal(IncomingAgentInfo.Integrity)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed Integrity encoding") {
		return
	}

	networkJson, err := json.Marshal(IncomingAgentInfo.Network)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed network encoding") {
		return
	}

	systemJson, err := json.Marshal(IncomingAgentInfo.System)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed system encoding") {
		return
	}

	TelemetryData := db.UpsertAgentTelemetryParams{
		AgentID: IncomingAgentInfo.AgentId,
		AgentToken: pgtype.Text{
			String: IncomingAgentInfo.AgentToken,
			Valid:  true,
		},
		MachineID: IncomingAgentInfo.MachineID,
		Column4:   systemJson,
		Column5:   securityJson,
		Column6:   processJson,
		Column7:   integrityJson,
		Column8:   networkJson,
	}
	fmt.Println(TelemetryData)
	err = h.DB.UpsertAgentTelemetry(ctx, TelemetryData)
	if utils.CheckError(w, err, http.StatusInternalServerError, "failed creating telemetry") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "telemetry created successfully",
	})
}

func (h *TelemetryType) GetLatestTelemetryData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "user not found",
		})
	}
	if utils.CheckError(w, err, http.StatusBadRequest, "internal server error") {
		return
	}

	AllAgetns, err := h.DB.GetAllAgentByUserId(ctx, int64(user.ID))
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "agent not found",
		})
		return
	}

	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errors") {
		return
	}

	TelemetryList := []any{}

	for _, agent := range AllAgetns {
		telemetry, err := h.DB.GetLatestTelemetry(ctx, agent.AgentID)
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		TelemetryList = append(TelemetryList, map[string]any{
			"agent_id":    telemetry.AgentID,
			"machine_id":  telemetry.MachineID,
			"agent_token": telemetry.AgentToken,
			"network":     json.RawMessage(telemetry.NetworkData),
			"system":      json.RawMessage(telemetry.SystemData),
			"security":    json.RawMessage(telemetry.SecurityData),
			"proccess":    json.RawMessage(telemetry.ProcessesData),
			"integrity":   json.RawMessage(telemetry.IntegrityData),
		})

	}

	if len(TelemetryList) == 0 {
		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "no telemetry data found!",
			"data":    []string{},
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "fetched telemetry Data successfully!",
		"data":    TelemetryList,
	})

}
