package clients

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func NewApiClient() *resty.Client {
	baseURL := "https://practice.expandtesting.com/notes/api"

	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Accept", "application/json").
		SetTimeout(10 * time.Second)

	return client
}
