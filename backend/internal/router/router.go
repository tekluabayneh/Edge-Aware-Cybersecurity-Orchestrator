package router

import (
	"encoding/json"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	OAuthconfig "github.com/edge-aware-cyberSecurity/internal/OAuthConfig"
	"github.com/edge-aware-cyberSecurity/internal/handler"
	middlewareGlobal "github.com/edge-aware-cyberSecurity/internal/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func LoadRouter(db *db.Queries) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middlewareGlobal.ErrorHandler)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "server response with 200 ok message",
		})
	})

	router.With(middlewareGlobal.LoginMiddleWare).Route("/auth/l/", func(route chi.Router) {
		AuthLogin(route, db)
	})
	router.With(middlewareGlobal.RegisterMiddleWare).Route("/auth/r/", func(route chi.Router) {
		AuthRegister(route, db)
	})

	oauthHandler := &handler.OAuthHandler{DB: db}
	router.Get("/oauth/google/callback", oauthHandler.GoogleCallbackHandler)
	router.Get("/oauth/github/callback", oauthHandler.GitHubCallbackHandler)
	router.Get("/oauth/google", OAuthconfig.GoogleLoginHandler)
	router.Get("/oauth/github", OAuthconfig.GitHubLoginHandler)

	// device paring route
	router.With().Route("/", func(route chi.Router) {
		DeviceParing(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/paringToken", func(route chi.Router) {
		GenerateToken(route, db)
	})

	router.With().Route("/Token", func(route chi.Router) {
		AcknowledgePairing(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/telementary", func(route chi.Router) {
		TelemetryReport(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/device", func(route chi.Router) {
		AlertReport(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/alerts", func(route chi.Router) {
		GetSingleAlertByAgentId(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/alert", func(route chi.Router) {
		UpdateAlertByID(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/newAlert", func(route chi.Router) {
		CreateNewAlert(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/update", func(route chi.Router) {
		UpdateAllAlertToRead(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/status", func(route chi.Router) {
		GetAllAlertStatus(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/delete", func(route chi.Router) {
		DeleteAlertById(route, db)
	})

	router.With(middlewareGlobal.Authorize(db)).Route("/create", func(route chi.Router) {
		CreateCommand(route, db)
	})

	router.With().Route("/fetch", func(route chi.Router) {
		FeatchCommand(route, db)
	})

	router.With().Route("/ack", func(route chi.Router) {
		AcknowledgeCommandExecutionHandle(route, db)
	})

	router.With().Route("/res", func(route chi.Router) {
		AcknowledgeCommandExecutionResponseHandle(route, db)
	})

	return router
}
