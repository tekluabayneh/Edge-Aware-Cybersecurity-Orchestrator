package router

import (
	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	AuthPath "github.com/edge-aware-cyberSecurity/internal/auth"
	"github.com/edge-aware-cyberSecurity/internal/handler"
	"github.com/edge-aware-cyberSecurity/internal/handler/otp"
	"github.com/go-chi/chi"
)

func AuthLogin(router chi.Router, db *db.Queries) {
	AuthHandlerRoute := &AuthPath.AuthLoginHandlerType{
		DB: db,
	}
	router.Post("/login", AuthHandlerRoute.Login)
}

func AuthRegister(router chi.Router, db *db.Queries) {
	AuthHandlerRegisterRoute := &AuthPath.AuthRegisterHandlerType{
		DB: db,
	}
	router.Post("/register", AuthHandlerRegisterRoute.Register)
}

func DeviceParing(router chi.Router, db *db.Queries) {
	DeviceParingHandler := &handler.DevicePairingType{DB: db}
	router.Post("/DeviceParing", DeviceParingHandler.DevicePairing)
}
func GenerateToken(router chi.Router, db *db.Queries) {
	DeviceParingTokenHandler := &handler.DevicePairingType{DB: db}
	router.Post("/token", DeviceParingTokenHandler.GenerateTokenHandler)
}

func AcknowledgePairing(router chi.Router, db *db.Queries) {
	DeviceAckParingTokenHandler := &handler.DevicePairingType{DB: db}
	router.Post("/ack", DeviceAckParingTokenHandler.AcknowledgePairing)
}

func TelemetryReport(router chi.Router, db *db.Queries) {
	TelemetryReportHandler := &handler.TelemetryType{DB: db}
	router.Post("/report", TelemetryReportHandler.ReceiveTelemetry)
}

func GetLatestTelemetryData(router chi.Router, db *db.Queries) {
	TelemetryReportHandler := &handler.TelemetryType{DB: db}
	router.Get("/telemetry", TelemetryReportHandler.GetLatestTelemetryData)
}

func AlertReport(router chi.Router, db *db.Queries) {
	AlertReportHandler := &handler.AlertType{DB: db}
	router.Get("/alert", AlertReportHandler.Alerts)
}

func CreateNewAlert(router chi.Router, db *db.Queries) {
	AlertReportHandler := &handler.AlertType{DB: db}
	router.Post("/create", AlertReportHandler.CreateAlert)
}

func GetSingleAlertByAgentId(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.AlertType{DB: db}
	router.Get("/id", SingleAlertHandler.GetAlertByAgentId)
}

func UpdateAlertByID(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.AlertType{DB: db}
	router.Post("/id", SingleAlertHandler.UpdateAlertsById)
}

func UpdateAllAlertToRead(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.AlertType{DB: db}
	router.Patch("/alert", SingleAlertHandler.UpdateAllAlerts)
}

func GetAllAlertStatus(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.AlertType{DB: db}
	router.Get("/status", SingleAlertHandler.GetAllAlertStatus)
}

func DeleteAlertById(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.AlertType{DB: db}
	router.Delete("/delete", SingleAlertHandler.DeleteAlertById)
}

func CreateCommand(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.CreateCommdnType{DB: db}
	router.Post("/create", SingleAlertHandler.CreateCommandHandler)
}

func FeatchCommand(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.CreateCommdnType{DB: db}
	router.Post("/fetch", SingleAlertHandler.FetchPendingCommandsHandler)
}

func AcknowledgeCommandExecutionHandle(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.CreateCommdnType{DB: db}
	router.Post("/ack", SingleAlertHandler.AcknowledgeCommandExecutionHandler)
}

func AcknowledgeCommandExecutionResponseHandle(router chi.Router, db *db.Queries) {
	SingleAlertHandler := &handler.CreateCommdnType{DB: db}
	router.Post("/ack", SingleAlertHandler.AcknowledgeCommandExecutionResponseHandler)
}

func UPdateUserProfile(router chi.Router, db *db.Queries) {
	UpdateUserProfileHandler := &handler.UserType{DB: db}
	router.Post("/profile", UpdateUserProfileHandler.UpdateUserInfo)
}

func GetUserProfile(router chi.Router, db *db.Queries) {
	GetUserProfileHandler := &handler.UserType{DB: db}
	router.Get("/profile", GetUserProfileHandler.GetUserProfileInfo)
}

func GetAllAgent(router chi.Router, db *db.Queries) {
	GetAllAgentHandler := &handler.AgentType{DB: db}
	router.Get("/agent", GetAllAgentHandler.GetAllUserAgents)
}

func Enable2FA(router chi.Router, db *db.Queries) {
	TowFaeHandler := &otp.TowFaType{DB: db}
	router.Post("/2fa", TowFaeHandler.Enable2FA)
}

func Confirm2FA(router chi.Router, db *db.Queries) {
	TowFaeHandler := &otp.TowFaType{DB: db}
	router.Post("/2fa", TowFaeHandler.Confirm2FA)
}

func Verify2FA(router chi.Router, db *db.Queries) {
	TowFaeHandler := &otp.TowFaType{DB: db}
	router.Post("/verify", TowFaeHandler.Verify2FA)
}
