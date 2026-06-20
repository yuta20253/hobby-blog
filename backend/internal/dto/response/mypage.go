package response

type MypageResponse struct {
	User  PostUserResponse `json:"user"`
	Posts []PostResponse   `json:"posts"`
}
