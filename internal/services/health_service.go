package services

import (
	"github.com/go-resty/resty/v2"
	"github.com/x-typ/ginkgo-e2e/internal/models/health"
	"github.com/x-typ/ginkgo-e2e/internal/routes"
)

func CheckHealth(client *resty.Client) (*resty.Response, error) {
	var healthResponse health.HealthResponse

	resp, err := client.R().
		SetResult(&healthResponse).
		Get(routes.HealthCheckEndpoint)

	return resp, err
}
