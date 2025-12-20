package security

import "fmt"

type SuspiciousProcess struct {
	PID    int32
	Name   string
	Reason string
}

func DetectMaliciousProcesses() []SuspiciousProcess {
	fmt.Println("DetectMaliciousProcesses")
	return []SuspiciousProcess{}
}
