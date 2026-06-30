package response

import domainUser "hobby-blog/internal/domain/user"

func NewAuthUserResponse(user *domainUser.User) AuthUserResponse {
	return AuthUserResponse{
		ID:    uint(user.ID),
		Name:  user.Name.String(),
		Email: user.Email.String(),
	}
}
