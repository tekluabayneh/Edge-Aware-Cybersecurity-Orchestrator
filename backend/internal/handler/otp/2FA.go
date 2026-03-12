package otp

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	"github.com/edge-aware-cyberSecurity/internal/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type TowFaType struct {
	DB *db.Queries
}

type incomignCodeType struct {
	Code string `json:"code"`
}

type incomingEnableType struct {
	Enable string `json:"enable"`
}

func (h *TowFaType) Enable2FA(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := ctx.Value("email").(string)

	var incoming incomingEnableType
	json.NewDecoder(r.Body).Decode(&incoming)

	user, err := h.DB.GetUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "no user found",
		})
		return
	}

	record, err := h.DB.Get2FAByUser(ctx, int64(user.ID))

	// FIRST TIME SETUP
	if errors.Is(err, sql.ErrNoRows) {

		secret, otpURL, qrPNG, err := utils.Generate2FASetup(email)
		if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
			return
		}

		params := db.Create2FAParams{
			UserID:   int64(user.ID),
			FaSecret: secret,
		}

		h.DB.Create2FA(ctx, params)

		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "2FA QR generated",
			"data": map[string]any{
				"otpUrl": otpURL,
				"qr":     qrPNG,
			},
		})
		return
	}

	// DISABLE REQUEST
	if incoming.Enable == "disable" {
		Updateparams := db.UpdateIs2FAEnabledParams{
			UserID: int64(user.ID),
			Isenabled: pgtype.Bool{
				Bool:  false,
				Valid: true,
			},
		}
		err := h.DB.UpdateIs2FAEnabled(ctx, Updateparams)
		if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
			return
		}
		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "2FA disabled",
		})
		return
	}

	// ALREADY ENABLED
	if !record.Isenabled.Bool && incoming.Enable == "enable" {
		Updateparams := db.UpdateIs2FAEnabledParams{
			UserID: int64(user.ID),
			Isenabled: pgtype.Bool{
				Bool:  true,
				Valid: true,
			},
		}
		err := h.DB.UpdateIs2FAEnabled(ctx, Updateparams)
		if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
			return
		}
		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "2FA enabled",
		})
		return
	}
}

func (h *TowFaType) Confirm2FA(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	var incomignCode incomignCodeType
	err := json.NewDecoder(r.Body).Decode(&incomignCode)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}
	user, err := h.DB.GetUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "no user found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}
	SekretkeyFromDb, err := h.DB.Get2FAByUser(ctx, int64(user.ID))
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "no 2Fa secret key found for this user",
		})
		return
	}

	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}

	isValid, err := utils.Verify2FACode(SekretkeyFromDb.FaSecret, incomignCode.Code)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
		return
	}

	if !isValid {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"message": "2Fa code does not match or it's expired!",
		})
		return
	}

	Updateparams := db.UpdateIs2FAEnabledParams{
		UserID: int64(user.ID),
	}
	err = h.DB.UpdateIs2FAEnabled(ctx, Updateparams)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "2FA code sent to email check our inbox",
	})
}

// this functio what alwayes get used to check 2fa
func (h *TowFaType) Verify2FA(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.Context().Value("email").(string)
	var incomignCode incomignCodeType
	err := json.NewDecoder(r.Body).Decode(&incomignCode)
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		return
	}

	user, err := h.DB.GetUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "no user found",
		})
		return
	}
	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		print("a", err)
		return
	}

	SekretkeyFromDb, err := h.DB.Get2FAByUser(ctx, int64(user.ID))
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"message": "no 2Fa secret key found for this user",
		})
		return
	}

	if utils.CheckError(w, err, http.StatusInternalServerError, "internal server error") {
		print("b", err)
		return
	}

	isValid, err := utils.Verify2FACode(SekretkeyFromDb.FaSecret, incomignCode.Code)

	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
		return
	}

	if !isValid {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]any{
			"message": "2Fa code does not match or it's expired!",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "2Fa check passsed successfully",
	})
}
