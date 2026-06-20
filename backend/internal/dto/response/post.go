package response

type PostResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	Content    string              `json:"content"`
	User       PostUserResponse    `json:"user"`
	Category   CategoryResponse    `json:"category"`
	MediaFiles []MediaFileResponse `json:"media_files"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PostUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
