package telemetry

import (
	handler "agent/internal"
	"agent/internal/commands"
	"agent/internal/telemetry/integrity"
	"agent/internal/telemetry/network"
	"agent/internal/telemetry/processes"
	"agent/internal/telemetry/security"
	"agent/internal/telemetry/system"
	"agent/internal/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

type Agent struct {
	DeviceName   string `json:"device_name"`
	AgentID      string `json:"agent_id"`
	AgentToken   string `json:"agent_token"`
	Email        string `json:"email"`
	MachineID    string `json:"machine_id"`
	AgentVersion string `json:"agent_version"`
	OS           string `json:"os"`
	Status       string `json:"status"`
	LastSeen     string `json:"last_seen"`
}
type TelemetryType struct {
	Email      string                      `json:"email"`
	AgentId    string                      `json:"agent_id"`
	AgentToken string                      `json:"agent_token"`
	MachineID  string                      `json:"machine_id"`
	System     system.GetSysInfotype       `json:"system"`
	Security   security.SecurityReport     `json:"security"`
	Network    network.NetworkSnapshot     `json:"network"`
	Processes  []processes.ProcInfo        `json:"processes"`
	Integrity  integrity.IntegritySnapshot `json:"integrity"`
}

func Telemetry() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// AnalizerBaseURL := os.Getenv("ANALIZER_BASE_URL")
	AnalizerBaseURL, err := utils.FetchEnv()
	if err != nil {
		fmt.Println("ANALIZER_BASE_URL not set")
	}
	////////////////////////////////////////////instade of env get it from backend ///////////////////////////////////////////////////////
	if AnalizerBaseURL.ANALIZER_BASE_URL == "" {
		fmt.Println("ANALIZER_BASE_URL not set")
		return
	}

	var wg sync.WaitGroup

	SystemInfo := make(chan system.GetSysInfotype)
	Security := make(chan security.SecurityReport)
	Network := make(chan network.NetworkSnapshot)
	Processes := make(chan []processes.ProcInfo)
	Integrity := make(chan integrity.IntegritySnapshot)

	MainTicker := time.NewTicker(1 * time.Minute)
	defer MainTicker.Stop()

	go func() {
		Ticker := time.NewTicker(10 * time.Hour)
		defer Ticker.Stop()

		// for the first time call it
		handler.GetJwt()
		for {
			select {
			case <-Ticker.C:
				// after ten hour reached call the function again
				handler.GetJwt()
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		commandTicker := time.NewTicker(60 * time.Second)
		defer commandTicker.Stop()

		for {
			select {
			case <-commandTicker.C:
				go commands.Commands()
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return

		case <-MainTicker.C:
			wg.Add(1)

			go func() {
				defer wg.Done()
				// this files must be run all the time or must be invocked all the time just after 20 minute
				// and it shold not be blocked and it shold run in ti's own process and should not block other proccess
				// and alsos after teh main telemetry started it shoild not also block them cuz this fuctin shold run while teh telemetry is also working it's own job or may be put it inside the the telemetry so it cannot be blocked
				go security.Security(Security)
				go network.Network(Network)
				go processes.Processes(Processes)
				go integrity.Integrity(Integrity)
				go system.System(SystemInfo)

				sysInfo := <-SystemInfo
				security := <-Security
				network := <-Network
				process := <-Processes
				integrity := <-Integrity
				path := filepath.Join("internal/register", "token.txt")

				content, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				var agent Agent
				err = json.Unmarshal(content, &agent)
				if err != nil {
					fmt.Println("Error parsing JSON:", err)
					return
				}

				// finally send Telemetry to analizer
				TelementoryPaylod := TelemetryType{
					Email:      agent.Email,
					AgentId:    agent.AgentID,
					AgentToken: agent.AgentToken,
					MachineID:  agent.MachineID,
					System:     sysInfo,
					Security:   security,
					Network:    network,
					Processes:  process,
					Integrity:  integrity,
				}

				jsonPayload, err := json.Marshal(TelementoryPaylod)
				if err != nil {
					fmt.Println("JSON MARSHAL ERROR:", err)
					return
				}
				client := &http.Client{Timeout: time.Second * 10}
				req, err := http.NewRequestWithContext(ctx, "POST", AnalizerBaseURL.ANALIZER_BASE_URL+"/rawTelementory", bytes.NewReader(jsonPayload))

				token, err := utils.Getjwt()
				if err != nil {
					token = ""
				}
				req.Header.Add("Authorization", "Bearer "+token)
				if err != nil {
					fmt.Println("HTTP REQUEST ERROR:", err)
					return
				}

				req.Header.Set("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					fmt.Println("HTTP SEND ERROR:", err)
					return
				}

				defer res.Body.Close()
				if res.StatusCode != http.StatusOK {
					fmt.Printf("Bad response: %d %s\n", res.StatusCode, res.Status)
				}
				time.Sleep(time.Second * 5)
			}()
		}
	}
}
