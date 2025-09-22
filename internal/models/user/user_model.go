package user

type UpdateUserProfileRequest struct {
	Name    *string `json:"name,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Company *string `json:"company,omitempty"`
}
