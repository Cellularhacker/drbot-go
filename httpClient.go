package drbot

import (
	"crypto/tls"
	"net/http"
	"time"
)

var drbotClient *Client

func initClient() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 1000
	t.MaxIdleConnsPerHost = 1000
	t.IdleConnTimeout = 900 * time.Second
	t.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	drbotClient = &Client{
		client: &http.Client{
			Timeout:   600 * time.Second,
			Transport: t,
		},
		maxRetries: 5,
		retryDelay: 1 * time.Second,
	}
}

func GetDrbotClient() *Client {
	return drbotClient
}
