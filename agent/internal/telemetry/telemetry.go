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
	"net/http"
	"os"
	"sync"
	"time"
)

type TelemetryType struct {
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
						fmt.Println("reached 5 minutes")
						go commands.Commands()
					}
				}()

				// Collect system info
				sysInfo := <-SystemInfo
				security := <-Security
				network := <-Network
				process := <-Processes
				integrity := <-Integrity

				// finally send Telemetry to analizer
				TelementoryPaylod := TelemetryType{
					AgetnId:    "",
					AgentToken: "",
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
