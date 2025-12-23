package network

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

func Network(ch chan bool) {
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
				// checkedIpdsCollectin
				//  connections monitoring collect
				// check susIp only after ~41 times in an hour
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
					lastRun = time.Now()
				}
				val := ConnectionsMonitoring()
				fmt.Printf("data %+v\n", val)
				time.Sleep(5 * time.Second)
				ch <- true
				defer wg.Done()
			}()
			wg.Wait()
		}
	}
}
