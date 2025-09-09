package services

import (
	"github.com/go-resty/resty/v2"
	"github.com/x-typ/ginkgo-e2e/internal/models/auth"
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

func GetUserProfile(client *resty.Client, token string) (*resty.Response, error) {
	var userProfileResponse auth.UserProfileInfoResponse

	resp, err := client.R().
		SetHeader("x-auth-token", token).
		SetResult(&userProfileResponse).
		Get(routes.GetUserInfoEndpoint)
	return resp, err
}
