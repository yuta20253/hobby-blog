package response

import "hobby-blog/internal/model"

func NewMypageResponse(
	user model.User,
	posts []model.Post,
) *MypageResponse {
	return &MypageResponse{
		User:  NewPostUserResponse(user),
		Posts: NewPostResponses(posts),
	}
}
