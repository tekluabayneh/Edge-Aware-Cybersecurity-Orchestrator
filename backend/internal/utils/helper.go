package utils

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Agent struct {
	UserID       int64              `json:"user_id"`
	AgentToken   string             `json:"agent_token"`
	AgentId      string             `json:"agent_id"`
	MachineID    string             `json:"machine_id"`
	AgentVersion pgtype.Text        `json:"agent_version"`
	Os           pgtype.Text        `json:"os"`
	Status       pgtype.Text        `json:"status"`
	LastSeen     pgtype.Timestamptz `json:"last_seen"`
}

func RequiredFieldsSet(a Agent) bool {
	fields := []any{a.Os, a.AgentToken, a.MachineID, a.LastSeen, a.AgentVersion, a.Status, a.AgentId, a.UserID}
	for _, f := range fields {
		if f == "" {
			return false
		}
	}
	return true

}
