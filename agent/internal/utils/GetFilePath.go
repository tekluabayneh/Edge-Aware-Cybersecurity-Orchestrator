package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetStoragePath() (string, string) {
	var baseDir string

	switch runtime.GOOS {
	case "windows":
		// Force Windows to use Program Files or AppData
		// "C:\Program Files\AgentOrchestrator"
		exe, _ := os.Executable()
		baseDir = filepath.Dir(exe)

	case "linux":
		// Standard Linux path for app data
		baseDir = "/var/lib/agent-orchestrator"
	default:
		// macOS/Unix hidden home folder
		home, _ := os.UserHomeDir()
		baseDir = filepath.Join(home, ".agent-orchestrator")
	}

	finalPath := filepath.Join(baseDir, "internal", "register")
	RegisterFolder := filepath.Join(baseDir, "register")

	// Create it if it's missing
	err := os.MkdirAll(finalPath, 0755)
	if err != nil {
		fmt.Println("err man nega", err)
		panic(err)
	}

	_, err = os.Create(RegisterFolder)
	if err != nil {
		fmt.Println("err man nega", err)
		panic(err)
	}

	return finalPath, RegisterFolder
}
