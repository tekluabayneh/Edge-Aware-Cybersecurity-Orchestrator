package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IncomingEnv struct {
	BASE_URL          string `json:"BASE_URL"`
	ANALIZER_BASE_URL string `json:"ANALIZER_BASE_URL"`
	API_KEY_OF_SUSIP  string `json:"API_KEY_OF_SUSIP"`
}

func FetchEnv() (IncomingEnv, error) {
	discoveryURL := "https://gist.githubusercontent.com/tekluabayneh/7b75001d3acbc0e1d8965495d0408676/raw/config.json"
	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Get(discoveryURL)
	if err != nil {
		return IncomingEnv{}, err
	}

	defer res.Body.Close()

	var IncomingEnvData IncomingEnv
	err = json.NewDecoder(res.Body).Decode(&IncomingEnvData)
	if err != nil {
		return IncomingEnv{}, err
	}

	fmt.Println("ANALIZER_BASE_URL", IncomingEnvData.ANALIZER_BASE_URL)
	fmt.Println("backend_base_url", IncomingEnvData.BASE_URL)

	return IncomingEnvData, nil
}
