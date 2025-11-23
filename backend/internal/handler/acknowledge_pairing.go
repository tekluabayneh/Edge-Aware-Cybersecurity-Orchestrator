package handler

import (
	"net/http"
	"time"
)

type Agent struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	MachineID    string    `json:"machine_id"`
	AgentVersion string    `json:"agent_version"`
	OS           string    `json:"os"`
	Status       string    `json:"status"`
	LastSeen     time.Time `json:"last_seen"`
}

func (h *DevicePairingType) AcknowledgePairing(w http.ResponseWriter, r *http.Request) {

}
