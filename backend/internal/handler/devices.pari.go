package handler

import (
	"fmt"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
)

type DevicePairingType struct {
	DB *db.Queries
}

func (h *DevicePairingType) DevicePairing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("devince paring handler test")
	// TODO:
	/*
	   Generate token
	    -- When user clicks “Add Device”.
	   Store token related to the user email
	    -- Save it in pairing_tokens table with expiry.
	   Limit number of active tokens per user (optional but recommended now)
	     -- If there are too many, delete old tokens before proceeding.
	   Logging (optional but early helps debugging)
	     -- Record who generated token & timestamp.
	   Handle refresh token request
	     -- If the user clicked refresh → generate new token, replace old one, update DB.
	   Check if incoming token is valid & pair the device with the user
	     -- When agent sends pairing request → validate token → link agent → create agent_id and agent_token.
	   Agent should send those credentials and store in agent table
	     -- Store agent_token, agent_id, machine_id, OS, version, etc.
	   Expiration
	     -- Automatically expire pairing token after a few minutes (or when used).
	*/

}
