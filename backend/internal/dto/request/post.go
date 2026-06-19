package request

import "hobby-blog/internal/model"

type SearchPostQuery struct {
	Title    string
	UserName string
	Category string
	Limit    int
	Offset   int
}

type CreatePostInput struct {
	Title      string
	Content    string
	CategoryID uint
	UserID     uint
}

type UpdatePostInput struct {
	ID         uint
	Title      string
	Content    string
	CategoryID uint
	UserID     uint
	Status     model.Status
}
