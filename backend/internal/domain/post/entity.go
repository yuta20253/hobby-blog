package post

import "hobby-blog/internal/domain/user"

type Post struct {
	ID uint
	UserID user.ID
	CategoryID uint
	Title string
	Content string
	Status Status
}
