package register

import (
	"bufio"
	"fmt"
	"os"
)

// user generate code user dashboard
// user isntall agent withe the generated code
// device make request to the device paring api
// api will return generated toke
// device will send acknowlaege with the device info, in success the agent will start to run

func Register() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please enter your registration token: ")
		token, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		token = token[:len(token)-1]

		if token == "" {
			fmt.Println("Token cannot be empty. Please try again.")
			continue
		}

		return token, nil
	}
}
