package register

import (
	"agent/internal/httpclient"
	"agent/internal/utils"
	"bufio"
	"context"
	"encoding/json"
	"fmt"

	"os"
	"strings"
	"time"
)

// user generate code user dashboard
// user isntall agent withe the generated code
// device make request to the device paring api
// api will return generated toke
// device will send acknowlaege with the device info, in success the agent will start to run

type DeviceInfo struct {
	DeviceName   string    `json:"device_name"`
	AgentID      string    `json:"agent_id"`
	AgentToken   string    `json:"agent_token"`
	MachineID    string    `json:"machine_id"`
	AgentVersion string    `json:"agent_version"`
	OS           string    `json:"os"`
	Status       string    `json:"status"`
	LastSeen     time.Time `json:"last_seen"`
}

type PairingResponse struct {
	AgentID    string `json:"agent_id"`
	AgentToken string `json:"agent_token"`
	Message    string `json:"message"`
}

type PairingAckResponse struct {
	Ack     bool   `json:"ack"`
	Message string `json:"message"`
}

func askForToken() ([]string, error) {
	var token, email string
	reader := bufio.NewReader(os.Stdin)

	for token == "" || email == "" {
		if token == "" {
			fmt.Println(utils.PromptBox.Render("Please copy your token from the user dashboard and enter it below:"))
			fmt.Println(utils.PromptBox.Render(" Token "))
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				return []string{}, err
			}
			token = strings.TrimSpace(input)
			if token == "" {
				fmt.Println(utils.ErrorBox.Render("Token cannot be empty."))
			}
		}

		if email == "" {
			fmt.Println(utils.PromptBox.Render("Please enter the email you used to register in the dashboard:"))
			fmt.Println(utils.PromptBox.Render(" Email "))
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				return []string{}, err
			}
			email = strings.TrimSpace(input)
			if email == "" {
				fmt.Println(utils.ErrorBox.Render("Email cannot be empty."))
			}
		}
	}

	return []string{token, email}, nil
}

func Register() bool {
	usrePaylod, err := askForToken()
	if err != nil {
		fmt.Println("Token cannot be empty. Please try again.")
	}

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		panic("base url is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	email := map[string]any{"email": usrePaylod[1]}

	res, err := httpclient.Fetch(ctx, baseUrl+"/DeviceParing", email, usrePaylod[0])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer res.Body.Close()
	var responsPaylod PairingResponse
	json.NewDecoder(res.Body).Decode(&responsPaylod)

	if res.StatusCode != 200 {
		fmt.Println(utils.ErrorBox.Render("message: =>", responsPaylod.Message))
		return false
	}

	// if the device paring success make device acknowlaege api call
	DeviceInfoPaylod := DeviceInfo{
		DeviceName:   "teklu_dev",
		AgentID:      responsPaylod.AgentID,
		AgentToken:   responsPaylod.AgentToken,
		MachineID:    "machine-004",
		AgentVersion: "1.4.2",
		Status:       "pending",
		LastSeen:     time.Now(),
		OS:           "window",
	}

	var DeviceInfoRespons PairingAckResponse
	if res.StatusCode == 200 {
		res, err := httpclient.InitiateParing(ctx, baseUrl+"/Token/ack", DeviceInfoPaylod)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Println("res res", res)
		if res.StatusCode != 200 && !DeviceInfoRespons.Ack {
			fmt.Println(utils.ErrorBox.Render(DeviceInfoRespons.Message))
			return false
		}

	}

	// make api call withe tht token
	// get the response

	fmt.Println(utils.SuccessBox.Render("✅ DeviceParing complete!"))
	return true

}
