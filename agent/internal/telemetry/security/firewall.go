package security

import "fmt"

type FirewallStatus struct {
	Enabled bool
}

func CheckFirewall() FirewallStatus {
	fmt.Println("CheckFirewall")
	return FirewallStatus{}
}
