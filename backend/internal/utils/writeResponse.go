package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, Msg any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Msg)
}
