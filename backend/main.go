package main

import (
	"github.com/edge-aware-cyberSecurity/cmd/server"
	handler "github.com/edge-aware-cyberSecurity/internal/OAuthConfig"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading env")
	// }

	// load env files after env loaded
	handler.InitOAuthConfigs()
	app := server.New()
	app.Start()
}
