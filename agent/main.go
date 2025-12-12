package main

import (
	"agent/cmd/app"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	app.App()
	fmt.Print("main work")
}
