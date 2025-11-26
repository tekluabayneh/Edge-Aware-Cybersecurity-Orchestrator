package handler

import (
	"fmt"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
)

type AlertType struct {
	DB *db.Queries
}

func (h *AlertType) Alert(w http.ResponseWriter, r *http.Request) {

	fmt.Println("test alert work")

}
