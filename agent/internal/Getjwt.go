package handler

import (
	"agent/internal/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type JwtType struct {
	jwt string
}

type userinfotype struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type storedUserInfoType struct {
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

func GetJwt() {
	ctx := context.TODO()
	client := http.Client{Timeout: 10 * time.Second}

	// read the register files so it can get user info
	path := filepath.Join("internal/register", "token.txt")
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("failed to read file %s: %v", path, err)
		return
	}

	var storeUserinfo storedUserInfoType
	if err := json.Unmarshal(content, &storeUserinfo); err != nil {
		log.Printf("failed to unmarshal stored user info: %v", err)
		return
	}

	jsondata, err := json.Marshal(storeUserinfo)
	if err != nil {
		log.Printf("failed to marshal user info: %v", err)
		return
	}

	data := bytes.NewReader(jsondata)

	// BASEURL := os.Getenv("BASE_URL")
	///////////////////////////////////////// fetch the url isntade of from env////////////////////////////////////////////////
	// if BASEURL == "" {
	// 	log.Println("BASE_URL environment variable is empty")
	// 	return
	// }

	baseurl, err := utils.FetchEnv()
	if err != nil {
		log.Printf("failed to request env file: %v", err)
	}

	url := baseurl.BASE_URL + "/api/auth/l/login"

	req, err := http.NewRequestWithContext(ctx, "POST", url, data)
	if err != nil {
		log.Printf("failed to create request: %v", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("request failed: %v", err)
		return
	}

	defer res.Body.Close()

	var userinfo userinfotype
	if err := json.NewDecoder(res.Body).Decode(&userinfo); err != nil {
		log.Printf("failed to decode response body: %v", err)
		return
	}

	if userinfo.Token == "" {
		log.Println("received empty token from server")
		return
	}

	// if the fiels does not exist creat the fiels and append it
	tokenpath := filepath.Join("internal/register", "Jwttoken.txt")
	_, err = os.Stat(tokenpath)
	if os.IsNotExist(err) {
		_, err := os.Create(tokenpath)
		if err != nil {
			fmt.Println("error creating Jwttoken file")
			return
		}
	} else if err != nil {
		fmt.Println("error creating Jwttoken file")
	}

	// check if it exist if not creat the file if does exist just use that

	if err := os.WriteFile(tokenpath, []byte(userinfo.Token), 0644); err != nil {
		log.Printf("failed to write token file %s: %v", tokenpath, err)
		return
	}

	log.Println("JWT token written successfully")
}
