package response

type AuthUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthResult struct {
	User  AuthUserResponse `json:"user"`
	Token string           `json:"token"`
}
