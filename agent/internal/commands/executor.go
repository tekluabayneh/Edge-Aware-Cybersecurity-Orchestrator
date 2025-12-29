package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type CommandType struct {
	IsAllCommandApproved bool `json:"isAllCommandApproved"`
}

type CommandItem struct {
	ID          int       `json:"ID"`
	UserID      int       `json:"UserID"`
	AgentID     string    `json:"AgentID"`
	CommandType string    `json:"CommandType"`
	Payload     string    `json:"Payload"`
	Status      string    `json:"Status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

type Response struct {
	Command []CommandItem     `json:"command"`
	Message string            `json:"message"`
	Payload map[string]string `json:"payload"`
}

// 1 agent recive commands from where the status is pending
// 2 the Commands payload has machineId and command to execute them
// 3 execute the command one by one if they are more than one
// 4 user can't send Commands that need admin to execute that Commands

func Commands(ch chan CommandType) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	URL := os.Getenv("BASE_URL")
	fullUrl := URL + "/fetch/fetch"
	path := filepath.Join("internal/register", "email.txt")

	if _, err := os.Stat(path); err != nil {
		fmt.Println("STAT ERROR:", err)
	}

	file, err := os.ReadFile(path)
	Bytesemail := strings.TrimSpace(string(file))
	if err != nil {
		panic(err)
	}
	Emailpayload := map[string]string{
		"email": Bytesemail,
	}
	jsonPayload, err := json.Marshal(Emailpayload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl, bytes.NewReader(jsonPayload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	var incomingPaylod Response
	json.NewDecoder(res.Body).Decode(&incomingPaylod)
	isAllCommandApproved := false
	for cmdName, args := range incomingPaylod.Payload {
		fmt.Printf("Executing command: %s %v\n", cmdName, args)
		cmd := exec.Command(cmdName, args)
		output, err := cmd.CombinedOutput()
		if err != nil {
			isAllCommandApproved = false
			fmt.Printf(" Command failed: %s %v\nError: %v\nOutput: %s\n", cmdName, args, err, string(output))
			continue
		} else {
			isAllCommandApproved = true
		}
	}

	fmt.Println(isAllCommandApproved)
}
