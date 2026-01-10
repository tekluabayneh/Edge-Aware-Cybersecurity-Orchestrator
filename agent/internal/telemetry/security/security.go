package security

import (
	"context"
	"sync"
	"time"
)

type SecurityReport struct {
	Firewall           FirewallStatus        `json:"firewall"`
	Antivirus          AntivirusStatus       `json:"antivirus"`
	MaliciousProcesses []SuspiciousProcess   `json:"malicious_processes"`
	SuspiciousFiles    []SuspiciouseFiletype `json:"suspicious_files"`
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
