package telemetry

import (
	"agent/internal/commands"
	"agent/internal/telemetry/integrity"
	"agent/internal/telemetry/network"
	"agent/internal/telemetry/processes"
	"agent/internal/telemetry/security"
	"agent/internal/telemetry/system"
	"context"
	"fmt"
	"sync"
	"time"
)

func Telemetry() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	chanSystemInfo := make(chan system.GetSysInfotype)
	chanSecurity := make(chan security.SecurityReport)
	chanNetwork := make(chan network.NetworkSnapshot)
	chanProcesses := make(chan []processes.ProcInfo)
	chanIntegrity := make(chan integrity.IntegritySnapshot)
	chanCommands := make(chan commands.CommandType)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			wg.Add(1)
			go func() {

				go security.Security(chanSecurity)
				go network.Network(chanNetwork)
				go processes.Processes(chanProcesses)
				go integrity.Integrity(chanIntegrity)
				go system.System(chanSystemInfo)
				go commands.Commands(chanCommands)

				// Collect system info
				// sysInfo := <-chanSystemInfo
				// securityr := <-chanSecurity
				// networkr := <-chanNetwork
				// processr := <-chanProcesses
				// integr := <-chanIntegrity
				commands := <-chanCommands
				//
				// fmt.Println("\n", sysInfo)
				// fmt.Println("\n", securityr)
				// fmt.Println("\n", networkr)
				// fmt.Println("\n", processr)
				// fmt.Println("\n", integr)

				// send comamnd to directly user dashboard
				fmt.Println("\n", commands)

				// send all alert to analizer
				defer wg.Done()
				time.Sleep(time.Second * 5)
			}()
			wg.Wait()
		}
	}
}
