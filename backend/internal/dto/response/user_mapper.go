package response

import "hobby-blog/internal/model"

func NewPostUserResponse(user model.User) PostUserResponse {
	return PostUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
