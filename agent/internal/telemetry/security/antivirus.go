package security

import (
	"os/exec"
	"runtime"
	"strings"
)

type AntivirusStatus struct {
	Running  bool   `json:"running"`
	Name     string `json:"name"`
	Detected string `json:"detected"`
}

func CheckAntivirus() AntivirusStatus {
	switch runtime.GOOS {
	case "linux":
		out, err := exec.Command("systemctl", "list-units", "--type=service").Output()
		if err == nil && strings.Contains(string(out), "clamav") {
			return AntivirusStatus{Running: true, Name: "clamav", Detected: "clamav detected"}
		} else {
			return AntivirusStatus{Running: false, Name: "unknown", Detected: "no known AV detected"}
		}
	case "darwin":
		out, err := exec.Command("ls", "/System/Library/CoreServices/XProtect.bundle").Output()
		if err == nil && len(out) > 0 {
			return AntivirusStatus{Running: true, Name: "", Detected: "built-in protection present"}
		}
		return AntivirusStatus{Running: false, Name: "unknown", Detected: "status unknown"}
	case "windows":
		return AntivirusStatus{Running: true, Name: "window defender"}
	default:
		return AntivirusStatus{Running: false, Name: "unsupported OS"}
	}
}
