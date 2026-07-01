package response

import (
	postInfrastructureModel "hobby-blog/internal/post/infrastructure"
	postPresentationResponse "hobby-blog/internal/post/presentation"
	userPresentationResponse "hobby-blog/internal/user/presentation"
)

func NewMypageResponse(
	user userPresentationResponse.AuthUserResponse,
	posts []postInfrastructureModel.Post,
) *MypageResponse {
	return &MypageResponse{
		User:  user,
		Posts: postPresentationResponse.NewPostResponses(posts),
	}
}
