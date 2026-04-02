package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"agent/internal/register"
	"agent/internal/telemetry"
	"agent/internal/utils"
)

// first the cli cick in
// the cli get validated if it contain the token
// the deviceparing api will be called then that will send the token match the token with the sotored token if the toke match  continue
// then acknowledge api will be alled witht agent info
// after succeess agent will send raw data to analizer
// loop will continue forever

func App() {
	// check registration status
	_, folderPath := utils.GetStoragePath()
	value, err := os.ReadFile(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("📁 Register file doesn't exist yet (first run) - proceeding to registration")
		} else {
			fmt.Println("🚨 READ FILE ERROR:", err)
		}
		value = []byte{}
	} else {
		fmt.Printf("✅ Read existing registration: %q\n", string(value))
	}
	if err != nil {
		fmt.Println("READ FILE ERROR:", err)
		value = []byte{}
	}
	time.Sleep(time.Second * 5)

	if len(value) < 1 || !strings.Contains(string(value), "registred") {
		// authenticate agent with user
		isRegistered := register.Register()
		if !isRegistered {
			fmt.Println("User not registered")
			return
		}

		err := utils.SafeWriteFile(folderPath, []byte("registred"))
		if err != nil {
			fmt.Println("WRITE FILE ERROR:", err)
			return
		}
	}

	// start the main app
	telemetry.Telemetry()
}
