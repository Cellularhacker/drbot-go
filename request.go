package drbot

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	PathAPI = "api"
)

func makeRequest(method, endpoint, path string, p url.Values, body []byte) ([]byte, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if p != nil {
		path = fmt.Sprintf("%s?%s", path, p.Encode())
	}

	req, _ := http.NewRequest(method, fmt.Sprintf("%s/%s", endpoint, path), bytes.NewBuffer(body))

	// MARK: clone from proxy2.DefaultClient
	client := GetDrbotClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to drbot api: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body from drbot api: %w", err)
	}

	return respBody, nil
}

func MakeRequest(method, path string, p url.Values, body []byte) ([]byte, error) {
	return makeRequest(method, apiEndpoint, path, p, body)
}

func MakeRequestChat(method, path string, p url.Values, body []byte) ([]byte, error) {
	return makeRequest(method, chatAPIEndpoint, path, p, body)
}

func MakeRequestAdmin(method, path string, p url.Values, body []byte) ([]byte, error) {
	return makeRequest(method, adminAPIEndpoint, path, p, body)
}
