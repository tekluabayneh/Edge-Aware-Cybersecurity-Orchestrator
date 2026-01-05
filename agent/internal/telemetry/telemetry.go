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
	AgetnId    string                      `json:"agent_id"`
	AgentToken string                      `json:"agent_token"`
	SystemInfo system.GetSysInfotype       `json:"system"`
	Security   security.SecurityReport     `json:"security"`
	Netwrok    network.NetworkSnapshot     `json:"network"`
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

				// Collect system info
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
					AgetnId:    agent.AgentID,
					AgentToken: agent.AgentToken,
					SystemInfo: sysInfo,
					Security:   security,
					Netwrok:    network,
					Processes:  process,
					Integrity:  integrity,
				}

				jsonpaylod, err := json.Marshal(TelementoryPaylod)

				if err != nil {
					fmt.Println("JSON MARSHAL ERROR:", err)
				}

				req, err := http.NewRequestWithContext(ctx, "POST", AnalizerBaseURL+"/rawTelementory", bytes.NewReader(jsonpaylod))
				if err != nil {
					fmt.Println("HTTP REQUEST ERROR:", err)
				}

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Println("HTTP DO ERROR:", err)
				}
				fmt.Println(res)
				defer wg.Done()
				time.Sleep(time.Second * 5)
			}()
			wg.Wait()
		}
	}
}
