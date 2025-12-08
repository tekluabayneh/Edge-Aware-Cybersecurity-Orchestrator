package telemetry

import (
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
	chanSecurity := make(chan bool)
	chanNetwork := make(chan bool)
	chanProcesses := make(chan bool)
	chanIntegrity := make(chan bool)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			go security.Security(chanSecurity)
			go network.Network(chanNetwork)
			go processes.Processes(chanProcesses)
			go integrity.Integrity(chanIntegrity)
			go system.System(chanSystemInfo)

			// Collect system info
			sysInfo := <-chanSystemInfo
			securityR := <-chanSecurity
			networkR := <-chanNetwork
			processR := <-chanProcesses
			integR := <-chanIntegrity

			fmt.Println(sysInfo.Cpu[0])
			fmt.Println(securityR)
			fmt.Println(networkR)
			fmt.Println(processR)
			fmt.Println(integR)

			// send all alert to analizer
			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(time.Second * 5)
				fmt.Println("pro pro")
			}()
			wg.Wait()
		}
	}
}
