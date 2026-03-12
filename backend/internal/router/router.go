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

	router.Route("/api", func(api chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "server response with 200 ok message",
			})
		})

		api.With(middlewareGlobal.LoginMiddleWare).Route("/auth/l/", func(route chi.Router) {
			AuthLogin(route, db)
		})
		api.With(middlewareGlobal.RegisterMiddleWare).Route("/auth/r/", func(route chi.Router) {
			AuthRegister(route, db)
		})

		oauthHandler := &handler.OAuthHandler{DB: db}
		api.Get("/oauth/google/callback", oauthHandler.GoogleCallbackHandler)
		api.Get("/oauth/github/callback", oauthHandler.GitHubCallbackHandler)
		api.Get("/oauth/google", OAuthconfig.GoogleLoginHandler)
		api.Get("/oauth/github", OAuthconfig.GitHubLoginHandler)

		// device paring route
		api.With().Route("/", func(route chi.Router) {
			DeviceParing(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/paringToken", func(route chi.Router) {
			GenerateToken(route, db)
		})

		api.With().Route("/Token", func(route chi.Router) {
			AcknowledgePairing(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/telementary", func(route chi.Router) {
			TelemetryReport(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/device", func(route chi.Router) {
			AlertReport(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/alerts", func(route chi.Router) {
			GetSingleAlertByAgentId(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/alert", func(route chi.Router) {
			UpdateAlertByID(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/newAlert", func(route chi.Router) {
			CreateNewAlert(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/update", func(route chi.Router) {
			UpdateAllAlertToRead(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/status", func(route chi.Router) {
			GetAllAlertStatus(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/delete", func(route chi.Router) {
			DeleteAlertById(route, db)
		})

		api.With(middlewareGlobal.Authorize(db)).Route("/create", func(route chi.Router) {
			CreateCommand(route, db)
		})

		api.With().Route("/fetch", func(route chi.Router) {
			FeatchCommand(route, db)
		})

		api.With().Route("/ack", func(route chi.Router) {
			AcknowledgeCommandExecutionHandle(route, db)
		})

		api.With().Route("/res", func(route chi.Router) {
			AcknowledgeCommandExecutionResponseHandle(route, db)
		})
	})

	return router
}
