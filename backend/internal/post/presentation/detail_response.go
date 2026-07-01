package presentation

type PostDetailResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	Content    string              `json:"content"`
	User       PostUserResponse    `json:"user"`
	Category   CategoryResponse    `json:"category"`
	MediaFiles []MediaFileResponse `json:"media_files"`
}
