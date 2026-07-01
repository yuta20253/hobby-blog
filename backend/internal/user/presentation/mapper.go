package presentation

import domainUser "hobby-blog/internal/user/domain"

func NewAuthUserResponse(user *domainUser.User) AuthUserResponse {
	return AuthUserResponse{
		ID:    uint(user.ID),
		Name:  user.Name.String(),
		Email: user.Email.String(),
	}
}
