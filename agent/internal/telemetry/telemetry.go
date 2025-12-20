package telemetry

import (
	// "agent/internal/telemetry/integrity"
	// "agent/internal/telemetry/network"
	// "agent/internal/telemetry/processes"
	"agent/internal/telemetry/security"
	// "agent/internal/telemetry/system"
	"context"
	"fmt"
	"sync"
	"time"
)

func Telemetry() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	// chanSystemInfo := make(chan system.GetSysInfotype)
	chanSecurity := make(chan security.SecurityReport)
	// chanNetwork := make(chan bool)
	// chanProcesses := make(chan bool)
	// chanIntegrity := make(chan bool)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			wg.Add(1)
			go func() {

				go security.Security(chanSecurity)
				// go network.Network(chanNetwork)
				// go processes.Processes(chanProcesses)
				// go integrity.Integrity(chanIntegrity)
				// go system.System(chanSystemInfo)

				// Collect system info
				// sysInfo := <-chanSystemInfo
				securityr := <-chanSecurity
				// networkr := <-chanNetwork
				// processr := <-chanProcesses
				// integR := <-chanIntegrity
				//
				// fmt.Println(sysInfo)
				fmt.Println(securityr)
				// fmt.Println(networkR)
				// fmt.Println(processr)
				// fmt.Println(integr)

				// send all alert to analizer
				defer wg.Done()
				time.Sleep(time.Second * 5)
			}()
			wg.Wait()
		}
	}
}
