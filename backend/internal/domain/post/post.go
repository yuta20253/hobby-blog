package post

type SearchQuery struct {
	Title    string
	UserName string
	Category string
	Limit    int
	Offset   int
}

type CreateInput struct {
	Title      string
	Content    string
	CategoryID uint
	UserID     uint
}

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
)

type UpdateInput struct {
	ID         uint
	Title      string
	Content    string
	CategoryID uint
	UserID     uint
	Status     Status
}
