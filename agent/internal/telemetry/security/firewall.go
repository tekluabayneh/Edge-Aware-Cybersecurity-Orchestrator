package security

import (
	"os/exec"
	"runtime"
	"strings"
)

type FirewallStatus struct {
	Enabled bool `json:"enabled"`
}

func CheckFirewall() FirewallStatus {
	switch runtime.GOOS {
	case "linux":
		out, err := exec.Command("systemctl", "status", "nftables").Output()
		if err == nil && strings.Contains(string(out), "Status: active") {
			return FirewallStatus{Enabled: true}
		}
		out, err = exec.Command("firewall-cmd", "--state").Output()
		if err == nil && strings.TrimSpace(string(out)) == "running" {
			return FirewallStatus{Enabled: true}
		}
		return FirewallStatus{Enabled: false}
	case "darwin":
		out, err := exec.Command("/usr/libexec/ApplicationFirewall/socketfilterfw", "--getglobalstate").Output()
		if err == nil && strings.Contains(string(out), "enabled") {
			return FirewallStatus{Enabled: true}
		}
		return FirewallStatus{Enabled: false}
	case "windows":
		out, err := exec.Command("netsh", "advfirewall", "show", "allprofiles", "state").Output()
		if err == nil && strings.Contains(string(out), "ON") || strings.Contains(string(out), "Enabled") {
			return FirewallStatus{Enabled: true}
		}
		return FirewallStatus{Enabled: false}
	default:
		return FirewallStatus{Enabled: false}
	}
}
