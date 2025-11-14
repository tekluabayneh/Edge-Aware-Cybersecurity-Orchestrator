package router

import (
	"encoding/json"
	"net/http"

	db "github.com/edge-aware-cyberSecurity/db/sqlc"
	AuthPath "github.com/edge-aware-cyberSecurity/internal/auth"
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

	return router
}

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
