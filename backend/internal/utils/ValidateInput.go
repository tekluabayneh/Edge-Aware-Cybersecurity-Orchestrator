package utils

import (
	"log"
	"regexp"
)

func IsINputSafe(value string) bool {
	pattern, err := regexp.Compile(`(?i)<script|<\/script|on\w+=|javascript:`)
	if err != nil {
		log.Fatal("failed to check malware")
	}
	return !pattern.MatchString(value)
}
