package security

import (
	"log"

	"github.com/shirou/gopsutil/v3/process"
)

type SuspiciousProcess struct {
	PID      int32
	Name     string
	Path     string
	UserName string
}

func DetectMaliciousProcesses() []SuspiciousProcess {
	procs, _ := process.Processes()
	var SuspiciousProcessData []SuspiciousProcess

	for _, proc := range procs {
		cmd, err := proc.Cmdline()
		if err != nil {
			log.Printf(
				"failed to get cmdline for pid %d: %v",
				proc.Pid,
				err,
			)
			continue
		}

		name, _ := proc.Name()
		username, _ := proc.Username()
		proccessInfo := SuspiciousProcess{
			Name:     name,
			PID:      proc.Pid,
			UserName: username,
			Path:     cmd,
		}
		SuspiciousProcessData = append(SuspiciousProcessData, proccessInfo)
	}
	return SuspiciousProcessData
}
