package telemetry

import (
	"agent/internal/commands"
	"agent/internal/telemetry/integrity"
	"agent/internal/telemetry/network"
	"agent/internal/telemetry/processes"
	"agent/internal/telemetry/security"
	"agent/internal/telemetry/system"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	System     system.GetSysInfotype       `json:"system"`
	Security   security.SecurityReport     `json:"security"`
	Network    network.NetworkSnapshot     `json:"network"`
	Processes  []processes.ProcInfo        `json:"processes"`
	Integrity  integrity.IntegritySnapshot `json:"integrity"`
}

func Telemetry() {
	ctx, cancel := context.WithCancel(context.Background())
	AnalizerBaseURL := os.Getenv("ANALIZER_BASE_URL")
	defer cancel()
	var wg sync.WaitGroup
	SystemInfo := make(chan system.GetSysInfotype)
	Security := make(chan security.SecurityReport)
	Network := make(chan network.NetworkSnapshot)
	Processes := make(chan []processes.ProcInfo)
	Integrity := make(chan integrity.IntegritySnapshot)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			wg.Add(1)
			go func() {

				go security.Security(Security)
				go network.Network(Network)
				go processes.Processes(Processes)
				go integrity.Integrity(Integrity)
				go system.System(SystemInfo)

				// execute comamnd periodically
				go func() {
					for {
						time.Sleep(time.Second * 60)
						go commands.Commands()
					}
				}()

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
					System:     sysInfo,
					Security:   security,
					Network:    network,
					Processes:  process,
					Integrity:  integrity,
				}

				payloadBytes, err := json.MarshalIndent(TelementoryPaylod, "", "  ")
				if err != nil {
					fmt.Println("Error marshaling payload:", err)
				} else {
					fmt.Println("=== FULL TELEMETRY PAYLOAD BEING SENT ===")
					fmt.Println(string(payloadBytes))
					fmt.Println("=======================================")
				}

				jsonPayload, err := json.Marshal(TelementoryPaylod)
				if err != nil {
					fmt.Println("JSON MARSHAL ERROR:", err)
					wg.Done()
					return
				}

				// time.Sleep(time.Minute * 6)
				client := &http.Client{Timeout: time.Second * 10}
				req, err := http.NewRequestWithContext(ctx, "POST", AnalizerBaseURL+"/rawTelementory", bytes.NewReader(jsonPayload))
				if err != nil {
					fmt.Println("HTTP REQUEST ERROR:", err)
					wg.Done()
					return
				}

				req.Header.Set("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					fmt.Println("HTTP SEND ERROR:", err)
					wg.Done()
					return
				}

				defer res.Body.Close()
				if res.StatusCode != http.StatusOK {
					fmt.Printf("Bad response: %d %s\n", res.StatusCode, res.Status)
				}

				wg.Done()
				time.Sleep(time.Second * 5)
			}()
			wg.Wait()
		}
	}
}
