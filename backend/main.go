package main

import (
	"log"

	"github.com/edge-aware-cyberSecurity/cmd/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env")
	}

	app := server.New()

	app.Start()
}
