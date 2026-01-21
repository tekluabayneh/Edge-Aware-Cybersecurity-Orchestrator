package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type AbuseIPDBResponse struct {
	IPAddress            string `json:"ipAddress"`
	AbuseConfidenceScore int    `json:"abuseConfidenceScore"`
	TotalReports         int    `json:"totalReports"`
	IsWhitelisted        bool   `json:"isWhitelisted"`
}

var (
	suspiciousIPs = make(map[string]struct{})
	abuseCache    = make(map[string]AbuseIPDBResponse)
	mu            sync.RWMutex
)

func CheckIsSusIp(ctx context.Context, ip string) (AbuseIPDBResponse, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	url := "https://api.abuseipdb.com/api/v2/check?ipAddress=" + ip + "&maxAgeInDays=90"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return AbuseIPDBResponse{}, err
	}

	req.Header.Set("Key", os.Getenv("API_KEY_OF_SUSIP"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "network-monitor")

	res, err := client.Do(req)
	if err != nil {
		return AbuseIPDBResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return AbuseIPDBResponse{}, fmt.Errorf("bad response: %s", res.Status)
	}

	var wrapper struct {
		Data AbuseIPDBResponse `json:"data"`
	}

	if err := json.NewDecoder(res.Body).Decode(&wrapper); err != nil {
		return AbuseIPDBResponse{}, err
	}

	return wrapper.Data, nil
}

func CollectLocalIPs() error {
	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	mu.Lock()
	defer mu.Unlock()

	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				suspiciousIPs[ipnet.IP.String()] = struct{}{}
			}
		}
	}

	return nil
}

func StartAbuseIPWorker(ctx context.Context) {
	ticker := time.NewTicker(20 * time.Minute)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				runAbuseChecks(ctx)
			}
		}
	}()
}

func runAbuseChecks(ctx context.Context) {
	mu.RLock()
	ips := make([]string, 0, len(suspiciousIPs))
	for ip := range suspiciousIPs {
		ips = append(ips, ip)
	}
	mu.RUnlock()

	for _, ip := range ips {
		resp, err := CheckIsSusIp(ctx, ip)
		if err != nil {
			fmt.Println("AbuseIPDB error:", err)
			continue
		}
		mu.Lock()
		abuseCache[ip] = resp
		mu.Unlock()
	}
}

func GetAbuseIPData() map[string]AbuseIPDBResponse {
	mu.RLock()
	defer mu.RUnlock()

	out := make(map[string]AbuseIPDBResponse)
	for k, v := range abuseCache {
		out[k] = v
	}
	return out
}
