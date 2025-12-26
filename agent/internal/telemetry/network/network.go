package network

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"
)

type ConnectionMonitoringType struct {
	ActiveSockets      []ActiveSocket
	NetworkInterfaces  []networkInterfaces
	ConnectionPatterns []ConnectionPatterns
}

type NetworkSnapshot struct {
	ConnectionMonitoring ConnectionMonitoringType
	AbuseIPDBResponse    map[string]AbuseIPDBResponse
}

func Network(ch chan NetworkSnapshot) {
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
				//  connections monitoring collect
				// check susIp only after 8 times in an hour
				// since the time expire use the existing sus ips data
				lastRun := time.Now()
				delay := time.Minute * 8
				if time.Since(lastRun) < delay.Abs() {
					susIp, err := FilterSusIp()
					if err != nil {
						log.Println("Warning: no suspicious IPs found or failed to fetch", err)
					}
					jsonValue, err := json.Marshal(susIp)
					if err != nil {
						log.Println("Warning: failed to marshal suspicious IPs:", err)
					}

					jsonData := map[string]AbuseIPDBResponse{}
					err = json.Unmarshal(jsonValue, &jsonData)
					if err != nil {
						log.Println("Warning: failed to unmarshal suspicious IPs:", err)
					}
					val := ConnectionsMonitoring()
					payload := NetworkSnapshot{
						ConnectionMonitoring: val,
						AbuseIPDBResponse:    jsonData,
					}
					lastRun = time.Now()
					ch <- payload
					time.Sleep(5 * time.Second)
					defer wg.Done()

				}
			}()
			wg.Wait()
		}
	}
}
