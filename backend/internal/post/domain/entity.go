package domain

import userDomain "hobby-blog/internal/user/domain"

type Post struct {
	ID         uint
	UserID     userDomain.ID
	CategoryID uint
	Title      string
	Content    string
	Status     Status
}
