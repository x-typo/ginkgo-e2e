package services

import (
	"github.com/go-resty/resty/v2"
	"github.com/x-typ/ginkgo-e2e/internal/models/login"
	"github.com/x-typ/ginkgo-e2e/internal/routes"
)

func LoginUser(client *resty.Client, email string, password string) (*resty.Response, error) {
	var loginResponse login.LoginResponse

	requestBody := map[string]string{
		"email":    email,
		"password": password,
	}

	resp, err := client.R().
		SetBody(requestBody).
		SetResult(&loginResponse).
		Post(routes.LoginEndpoint)

	return resp, err
}
