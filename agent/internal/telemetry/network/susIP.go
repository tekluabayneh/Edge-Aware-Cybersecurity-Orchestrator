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

// get the all machine ips
// check them against suspiouse ip if so report them as suspiouse
// make api call to suspiouse ip and get the ips
// if since the cota of the ip is 100 a day this function only need to call this api 100 time with in 24hr
// besically i can say this function only need to called

type AbuseIPDBResponse struct {
	Data struct {
		IPAddress            string `json:"ipAddress"`
		AbuseConfidenceScore int    `json:"abuseConfidenceScore"`
		TotalReports         int    `json:"totalReports"`
		IsWhitelisted        bool   `json:"isWhitelisted"`
	} `json:"data"`
}

var (
	collectSusIp          = make(map[string]string)
	ipsCollectionToReturn = make(map[string]AbuseIPDBResponse)
	mu                    sync.Mutex
)

func CheckIsSusIp(ip string) (map[string]AbuseIPDBResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// sleep for one second befoere making another api call since ips may be more than one
	ipString := string(ip)

	client := &http.Client{Timeout: time.Second * 10}
	FullUrl := "https://api.abuseipdb.com/api/v2/check?ipAddress=" + ipString + "&maxAgeInDays=90"
	API_KEY := os.Getenv("API_KEY_OF_SUSIP")

	req, err := http.NewRequestWithContext(ctx, "GET", FullUrl, nil)
	req.Header.Set("Key", API_KEY)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, fmt.Errorf("failed to create request for IP check: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to AbuseIPDB: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response from AbuseIPDB: %d %s", res.StatusCode, res.Status)
	}

	defer res.Body.Close()
	collectIpStatus := map[string]AbuseIPDBResponse{}

	var responseValue AbuseIPDBResponse
	err = json.NewDecoder(res.Body).Decode(&responseValue)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response from AbuseIPDB: %w", err)
	}

	collectIpStatus[ip] = responseValue
	return collectIpStatus, nil
}

// first code need to only run 5 times a day
// second it shouln't be blocking
// theird it should handle faliour properly
// four it should only execute those ips that is found in the collectSusIp and wait for the sleep time
func StartSusIpCheck(ctx context.Context) {
	const checksPerDay = 5
	interval := 24 * time.Hour / checksPerDay
	ticker := time.NewTicker(interval)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				currentIps := make([]string, 0, len(collectSusIp))
				mu.Lock()
				for ip := range collectSusIp {
					currentIps = append(currentIps, ip)
				}
				mu.Unlock()

				if len(currentIps) == 0 {
					fmt.Println("No suspicious IPs to check at", time.Now().Format(time.RFC3339))
					continue
				}
				fmt.Printf("Starting batch check #%d of day — %d suspicious IPs → %s\n", checksPerDay, len(currentIps), time.Now().Format("15:04:05"))
				for _, ip := range currentIps {
					chechedIP, err := CheckIsSusIp(ip)
					if err != nil {
						fmt.Println("failed to check suspicious IP")
						continue
					}
					mu.Lock()
					ipsCollectionToReturn[ip] = chechedIP[ip]
					mu.Unlock()
				}
			}
		}
	}()

}

func FilterSusIp() (map[string]AbuseIPDBResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	iface, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to list network interfaces: %w", err)
	}

	var ip net.IP
	for _, i := range iface {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, fmt.Errorf("failed to get addresses for interface %s: %w", i.Name, err)
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil {
				continue
			}
			ipstr := ip.String()
			collectSusIp[ipstr] = ipstr
		}
		StartSusIpCheck(ctx)
	}
	return ipsCollectionToReturn, nil
}
