package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type IncomingEnv struct {
	BASE_URL          string `json:"BASE_URL"`
	ANALIZER_BASE_URL string `json:"ANALIZER_BASE_URL"`
	API_KEY_OF_SUSIP  string `json:"API_KEY_OF_SUSIP"`
}

func FetchEnv() (IncomingEnv, error) {
	discoveryURL := "https://gist.githubusercontent.com/tekluabayneh/7b75001d3acbc0e1d8965495d0408676/raw/9a60a31d37d0a43c41c37542c52aad0e9afea013/config.json"
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

	return IncomingEnvData, nil
}
