package security

import (
	"context"
	"sync"
	"time"
)

type SecurityReport struct {
	Firewall           FirewallStatus
	Antivirus          AntivirusStatus
	MaliciousProcesses []SuspiciousProcess
	SuspiciousFiles    []SuspiciouseFiletype
}

func Security(ch chan SecurityReport) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	for {
		select {
		case <-ctx.Done():
			return
		default:
			wg.Add(1)
			go func() {
				time.Sleep(5 * time.Second)
				ch <- SecurityReport{
					Firewall:           CheckFirewall(),
					Antivirus:          CheckAntivirus(),
					MaliciousProcesses: DetectMaliciousProcesses(),
					SuspiciousFiles:    DetectSuspiciousFiles(),
				}
				wg.Done()
			}()
			wg.Wait()
		}
	}

}
