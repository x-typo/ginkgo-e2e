package login

type UserData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
type LoginResponse struct {
	Success bool     `json:"success"`
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}
