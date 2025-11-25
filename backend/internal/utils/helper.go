package utils

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type Agent struct {
	UserID       int64              `json:"user_id"`
	AgentToken   string             `json:"agent_token"`
	AgentId      int64              `json:"agent_id"`
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

func CheckError(w http.ResponseWriter, err error, status int, Msg string) {
	if err != nil {
		WriteJSON(w, status, map[string]string{
			"message": Msg,
		})
		return
	}
}
