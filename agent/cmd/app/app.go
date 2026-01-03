package app

import (
	"agent/internal/register"
	"agent/internal/telemetry"
	"fmt"
	"os"
	"strings"
)

// first the cli cick in
// the cli get validated if it contain the token
// the deviceparing api will be called then that will send the token match the token with the sotored token if the toke match  continue
// then acknowledge api will be alled witht agent info
// after succeess agent will send raw data to analizer
// loop will continue forever

func App() {
	// check registration status
	value, err := os.ReadFile("./register")
	if err != nil {
		fmt.Println("READ FILE ERROR:", err)
		value = []byte{}
	}

	if len(value) < 1 || !strings.Contains(string(value), "registred") {
		// authenticate agent with user
		isRegistered := register.Register()
		if !isRegistered {
			fmt.Println("User not registered")
			return
		}

		err := os.WriteFile("./register", []byte("registred"), 0644)
		if err != nil {
			fmt.Println("WRITE FILE ERROR:", err)
			return
		}
	}

	// start the main app
	telemetry.Telemetry()
}
