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
