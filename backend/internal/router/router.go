package router

import (
	"encoding/json"
	"net/http"

	AuthPath "github.com/edge-aware-cyberSecurity/internal/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)


func LoadRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
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

	router.Route("/auth", func(route chi.Router) {
		AuthLogin(route)
		AuthRegister(route)
	})

	return router
}


func AuthLogin(router chi.Router) {
	AuthHandlerRoute := &AuthPath.AuthLoginHandlerType{}

	router.Post("/login", AuthHandlerRoute.Login)
}



func AuthRegister(router chi.Router) {
	AuthHandlerRegisterRoute := &AuthPath.AuthRegisterHandlerType{}

	router.Post("/register", AuthHandlerRegisterRoute.Register)
}



 





