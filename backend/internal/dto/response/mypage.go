package response

import (
	postPresentationResponse "hobby-blog/internal/post/presentation"
	userPresentationResponse "hobby-blog/internal/user/presentation"
)

type MypageResponse struct {
	User  userPresentationResponse.AuthUserResponse `json:"user"`
	Posts []postPresentationResponse.PostResponse   `json:"posts"`
}
