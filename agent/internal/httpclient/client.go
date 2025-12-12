package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

func Fetch(ctx context.Context, Url string, email map[string]any, token string) (*http.Response, error) {
	params := url.Values{}
	params.Add("token", token)
	fullURL := Url + "?" + params.Encode()

	jsonpaylod, _ := json.Marshal(email)
	bodyToSend := bytes.NewReader(jsonpaylod)
	req, err := http.NewRequestWithContext(ctx, "POST", fullURL, bodyToSend)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return res, nil
}

func InitiateParing(ctx context.Context, Url string, body any) (*http.Response, error) {
	jsonpaylod, _ := json.Marshal(body)
	bodyToSend := bytes.NewReader(jsonpaylod)
	req, err := http.NewRequestWithContext(ctx, "POST", Url, bodyToSend)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
