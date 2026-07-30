package main

import (
	"github.com/edge-aware-cyberSecurity/cmd/server"
	handler "github.com/edge-aware-cyberSecurity/internal/OAuthConfig"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// log.Println(".env file not found, using environment variables")
	// }

	// load env files after env loaded
	handler.InitOAuthConfigs()
	app := server.New()
	app.Start()
}
