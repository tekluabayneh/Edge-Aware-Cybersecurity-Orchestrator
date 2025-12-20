package security

import "fmt"

type AntivirusStatus struct {
	Running bool
	Name    string
}

func CheckAntivirus() AntivirusStatus {
	fmt.Println("CheckAntivirus")
	return AntivirusStatus{}
}
