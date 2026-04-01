package notification

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
)

type NotificationType struct {
	DB *db.Queries
}

type incomignNotificationDataType struct {
	Id int `json:"id"`
}

func (h *NotificationType) GetAllNotification(w http.ResponseWriter, r *http.Request) {
	// get the user email and get the notification id and mark is as read
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := r.Context().Value("email").(string)
	var incomignNotificationData incomignNotificationDataType
	err := json.NewDecoder(r.Body).Decode(&incomignNotificationData)

	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}

	AllNotificaion, err := h.DB.GetAllNotifications(ctx, int64(user.ID))
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "you don't notification yet!",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errors") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message":      "notification updated",
		"notification": AllNotificaion,
	})
}

func (h *NotificationType) UpdateSingleNotification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	var incomignNotificationData incomignNotificationDataType
	err := json.NewDecoder(r.Body).Decode(&incomignNotificationData)

	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}

	params := db.GetNotificationByIdParams{
		ID:     int64(incomignNotificationData.Id),
		UserID: int64(user.ID),
	}

	fmt.Println("id", incomignNotificationData.Id)
	fmt.Println("userid", user.ID)
	_, err = h.DB.GetNotificationById(ctx, params)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "notification not found!",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errors") {
		return
	}

	UPdatnotificationParams := db.UpdateNotificationByIdParams{
		UserID: int64(user.ID),
		ID:     int64(incomignNotificationData.Id),
	}

	err = h.DB.UpdateNotificationById(ctx, UPdatnotificationParams)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errors") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "notification updated",
	})

}

func (h *NotificationType) UpdateAllNotification(w http.ResponseWriter, r *http.Request) {
	// get ser email maks all notoficaton as read
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	user, err := h.DB.GetUserByEmail(ctx, email)
	if utils.CheckError(w, err, http.StatusBadRequest, "user not found") {
		return
	}

	_, err = h.DB.GetAllAgentByUserId(ctx, int64(user.ID))
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "you don't notification yet!",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errors") {
		return
	}

	err = h.DB.UpdateAllNotificationByUserId(ctx, int64(user.ID))
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server errors") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "all notification updated",
	})

}
