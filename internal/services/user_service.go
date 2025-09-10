package services

import (
	"github.com/go-resty/resty/v2"
	"github.com/x-typ/ginkgo-e2e/internal/models/auth"
	"github.com/x-typ/ginkgo-e2e/internal/models/user"
	"github.com/x-typ/ginkgo-e2e/internal/routes"
)

func GetUserProfile(client *resty.Client, token string) (*resty.Response, error) {
	var userProfileResponse auth.UserProfileInfoResponse

	resp, err := client.R().
		SetHeader("x-auth-token", token).
		SetResult(&userProfileResponse).
		Get(routes.GetUserInfoEndpoint)
	return resp, err
}

func UpdateUserProfile(client *resty.Client, token string, payload user.UpdateUserProfileRequest) (*resty.Response, error) {
	var userProfileResponse auth.UserProfileInfoResponse

	resp, err := client.R().
		SetHeader("x-auth-token", token).
		SetBody(payload).
		SetResult(&userProfileResponse).
		Patch(routes.UpdateUserInfoEndpoint)

	return resp, err
}
