package auth

import "github.com/x-typ/ginkgo-e2e/internal/models/shared"

type LoginResponse struct {
	shared.BaseResponse
	Data UserData `json:"data"`
}

type UserProfileInfoResponse struct {
	shared.BaseResponse
	Data UserProfileData `json:"data"`
}
type UserData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserProfileData struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
}
