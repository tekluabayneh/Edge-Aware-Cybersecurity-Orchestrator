package network

import (
	"context"
	"log"
	"time"
)

type ConnectionMonitoringType struct {
	ActiveSockets      []ActiveSocket       `json:"ActiveSockets"`
	NetworkInterfaces  []networkInterfaces  `json:"NetworkInterfaces"`
	ConnectionPatterns []ConnectionPatterns `json:"ConnectionPatterns"`
}
type NetworkSnapshot struct {
	ConnectionMonitoring ConnectionMonitoringType     `json:"ConnectionMonitoring"`
	AbuseIPDBResponse    map[string]AbuseIPDBResponse `json:"AbuseIPDBResponse"`
}

func Network(ch chan NetworkSnapshot) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	CollectLocalIPs()
	StartAbuseIPWorker(ctx)
	var lastRun time.Time
	delay := time.Minute * 7

	for range ticker.C {
		// only refresh AbuseIPDB after delay
		if time.Since(lastRun) < delay {
			log.Println("Skipping AbuseIPDB fetch (time window not expired)")
			continue
		}
		susIp := GetAbuseIPData()
		ConnectionsMonitoringData := ConnectionsMonitoring()

		payload := NetworkSnapshot{
			ConnectionMonitoring: ConnectionsMonitoringData,
			AbuseIPDBResponse:    susIp,
		}

		lastRun = time.Now()
		ch <- payload
	}
}
