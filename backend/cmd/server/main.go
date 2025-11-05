package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edge-aware-cyberSecurity/internal/router"
)

type App struct {
	router http.Handler
}

func New() *App {

	app := &App{
		router: router.LoadRouter(),
	}

	return app

}
func (app *App) Start() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		json.Marshal(map[string]string{
			"message": "failed to load env file",
		})
	}

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: app.router,
	}

	fmt.Printf("Server is running on post %v", PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to run ", err.Error())
	}
}
