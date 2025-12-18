package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

// get the machine all ips
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

func CheckIsSusIp(ip string) (map[string]AbuseIPDBResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// sleep for one second befoere making another api call since ips may be more than one
	time.Sleep(time.Second * 1)
	ipString := string(ip)
	FullUrl := "https://api.abuseipdb.com/api/v2/check?ipAddress=" + ipString + "&maxAgeInDays=90"
	API_KEY := os.Getenv("API_KEY_OF_SUSIP")
	req, err := http.NewRequestWithContext(ctx, "GET", FullUrl, nil)
	req.Header.Set("Key", API_KEY)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return nil, fmt.Errorf("failed to create request for IP check: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to AbuseIPDB: %w", err)
	}

	collectIpStatus := map[string]AbuseIPDBResponse{}
	var responseValue AbuseIPDBResponse
	err = json.NewDecoder(res.Body).Decode(&responseValue)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response from AbuseIPDB: %w", err)
	}

	collectIpStatus[ip] = responseValue
	return collectIpStatus, nil
}

func FilterSusIp() (map[string]AbuseIPDBResponse, error) {

	collectSusIp := map[string]string{}
	ipsCollectionToReturn := map[string]AbuseIPDBResponse{}
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

		// call the api based on the ips number/length
		for _, ip := range collectSusIp {
			checkedIps, err := CheckIsSusIp(collectSusIp[ip])
			fmt.Println("man", checkedIps)
			if err != nil {
				return nil, fmt.Errorf("failed to check suspicious IP %s: %w", ip, err)
			}
			ipsCollectionToReturn[ip] = checkedIps[ip]
		}
	}
	return ipsCollectionToReturn, nil
}
