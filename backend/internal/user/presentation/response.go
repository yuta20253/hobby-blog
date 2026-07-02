package presentation

type AuthUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthResponse struct {
	User  AuthUserResponse `json:"user"`
	Token string           `json:"token"`
}
