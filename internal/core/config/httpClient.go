package config

import (
	"net/http"
	"time"
)

func CreateHTTPClient() *http.Client {
	client := http.Client{Timeout: 10 * time.Second}
	return &client
}
