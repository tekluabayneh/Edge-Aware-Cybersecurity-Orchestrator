package register

import (
	"agent/internal/httpclient"
	"agent/internal/utils"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"os"
	"strings"
	"time"
)

// user generate code from user dashboard
// user isntall agent
// and past the token intot he terminal input box and also email
// device make request to the device paring api
// api will return generated toke
// device will send acknowlaege with the device info, in success the agent will start to run

type DeviceInfo struct {
	DeviceName   string    `json:"device_name"`
	AgentID      string    `json:"agent_id"`
	AgentToken   string    `json:"agent_token"`
	Email        string    `json:"email"`
	MachineID    string    `json:"machine_id"`
	AgentVersion string    `json:"agent_version"`
	OS           string    `json:"os"`
	Status       string    `json:"status"`
	LastSeen     time.Time `json:"last_seen"`
}

type DeviceInfoStore struct {
	DeviceName   string    `json:"device_name"`
	AgentID      string    `json:"agent_id"`
	AgentToken   string    `json:"agent_token"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
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
	var token, password, email string
	reader := bufio.NewReader(os.Stdin)

	for token == "" || email == "" || password == "" {
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

		if password == "" {
			fmt.Println(utils.PromptBox.Render("Please enter the password you used to register in the dashboard:"))
			fmt.Println(utils.PromptBox.Render(" password "))
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				return []string{}, err
			}
			password = strings.TrimSpace(input)
			if password == "" {
				fmt.Println(utils.ErrorBox.Render("password cannot be empty."))
			}
		}

	}

	return []string{token, email, password}, nil
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
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
		DeviceName:   utils.StaticSysInfo().HostName,
		Email:        usrePaylod[1],
		AgentID:      responsPaylod.AgentID,
		AgentToken:   responsPaylod.AgentToken,
		MachineID:    utils.StaticSysInfo().MachineID,
		AgentVersion: utils.StaticSysInfo().AgentVersion,
		Status:       utils.StaticSysInfo().Status,
		LastSeen:     time.Now(),
		OS:           utils.StaticSysInfo().OS,
	}

	var DeviceInfoRespons PairingAckResponse
	// if the token vlidation api response ok, proccede to the token acknowlaege api call
	if res.StatusCode == 200 {
		res, err := httpclient.InitiateParing(ctx, baseUrl+"/Token/ack", DeviceInfoPaylod)
		print(err)
		if err != nil {
			fmt.Println(utils.ErrorBox.Render("message: =>", DeviceInfoRespons.Message))
			fmt.Println("ack: =>", DeviceInfoRespons.Ack)
			fmt.Println(err)
		}

		// check if the akc tokne api call is response is not StatusOK
		if res.StatusCode != 200 && !DeviceInfoRespons.Ack {
			fmt.Println(utils.ErrorBox.Render(DeviceInfoRespons.Message))
			fmt.Println(res)
			return false
		}
	}

	// if the device paring success make device acknowlaege api call
	store := DeviceInfoStore{
		DeviceName:   utils.StaticSysInfo().HostName,
		Email:        usrePaylod[1],
		Password:     usrePaylod[2],
		AgentID:      responsPaylod.AgentID,
		AgentToken:   responsPaylod.AgentToken,
		MachineID:    utils.StaticSysInfo().MachineID,
		AgentVersion: utils.StaticSysInfo().AgentVersion,
		Status:       utils.StaticSysInfo().Status,
		LastSeen:     time.Now(),
		OS:           utils.StaticSysInfo().OS,
	}

	path := filepath.Join("internal/register", "email.txt")
	err = os.WriteFile(path, []byte(DeviceInfoPaylod.Email), 0644)
	if err != nil {
		fmt.Println(utils.ErrorBox.Render("message: =>", "error storing usre email"))
	}

	path = filepath.Join("internal/register", "token.txt")
	json, err := json.Marshal(store)
	if err != nil {
		fmt.Println(utils.ErrorBox.Render("message: =>", "error while marshling struct"))
	}

	err = os.WriteFile(path, []byte(json), 0644)
	if err != nil {
		fmt.Println(utils.ErrorBox.Render("message: =>", "error storing usre email"))
	}

	fmt.Println(utils.SuccessBox.Render("✅ DeviceParing complete!"))
	return true

}
