package app

import (
	"agent/internal/register"
	"agent/internal/telemetry"
	"fmt"
)

// first the cli cick in
// the cli get validated if it contain the token
// the deviceparing api will be called then that will send the token match the token with the sotored token if the toke match  continue
// then acknowledge api will be alled witht agent info
// after succeess agent will send raw data to analizer
// loop will continue forever
func App() {

	// authenticate agent with user
	isRegister := register.Register()
	if !isRegister {
		fmt.Println("user not registred")
		return
	}

	// if register pass start the app
	telemetry.Telemetry()
}
