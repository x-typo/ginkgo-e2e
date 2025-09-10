package services

import (
	"github.com/go-resty/resty/v2"
	"github.com/x-typ/ginkgo-e2e/internal/models/auth"
	"github.com/x-typ/ginkgo-e2e/internal/models/shared"
	"github.com/x-typ/ginkgo-e2e/internal/routes"
)

func LoginUser(client *resty.Client, email string, password string) (*resty.Response, error) {
	var loginResponse auth.LoginResponse

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

func LogoutUser(client *resty.Client, token string) (*resty.Response, error) {
	var logoutResponse shared.BaseResponse

	resp, err := client.R().
		SetHeader("x-auth-token", token).
		SetResult(&logoutResponse).
		Delete(routes.LogoutEndpoint)
	return resp, err
}
