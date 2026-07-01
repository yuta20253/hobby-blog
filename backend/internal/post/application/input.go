package application

import postDomain "hobby-blog/internal/post/domain"

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
	Status     postDomain.Status
}
