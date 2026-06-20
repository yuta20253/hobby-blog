package response

import "hobby-blog/internal/model"

func NewAuthUserResponse(user model.User) AuthUserResponse {
	return AuthUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
