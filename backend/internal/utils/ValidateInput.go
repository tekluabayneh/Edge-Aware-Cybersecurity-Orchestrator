package utils

import (
	"log"
	"regexp"
)

func IsINputSafe(value string) bool {
	pattern, err := regexp.Compile(`(?i)(<\s*script\b|</\s*script\s*>|on\w+\s*=|javascript:|vbscript:|data:|<\s*iframe\b|<\s*object\b|<\s*embed\b|src\s*=\s*["']\s*javascript:|expression\s*\(|<\s*link\b|<\s*meta\b)`)
	if err != nil {
		log.Fatal("failed to check malware")
	}
	return !pattern.MatchString(value)
}
