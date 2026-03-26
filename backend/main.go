package main

import (
	// "log"

	"github.com/edge-aware-cyberSecurity/cmd/server"
	handler "github.com/edge-aware-cyberSecurity/internal/OAuthConfig"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading env")
	// }
	//
	// load env files after env loaded
	handler.InitOAuthConfigs()
	app := server.New()
	app.Start()
}
