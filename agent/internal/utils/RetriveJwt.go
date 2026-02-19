package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type tokentype string

func Getjwt() (string, error) {
	tokenpath := filepath.Join("internal/register", "Jwttoken.txt")
	content, err := os.ReadFile(tokenpath)
	if err != nil {
		log.Printf("failed to write token file %s: %v", tokenpath, err)
		return "", errors.New("error while reading file")
	}
	//
	token := strings.TrimSpace(string(content))
	if token == "" {
		log.Println("received empty token from server")
		return "", errors.New("error while reading file")
	}
	return token, nil
}
