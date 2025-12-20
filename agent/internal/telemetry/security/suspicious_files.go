package security

import "fmt"

type SuspiciousFile struct {
	Path   string
	Reason string
}

func DetectSuspiciousFiles() []SuspiciousFile {
	fmt.Println("DetectSuspiciousFiles")
	return []SuspiciousFile{}
}
