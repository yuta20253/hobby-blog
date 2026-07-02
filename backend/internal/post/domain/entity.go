package domain

import userDomain "hobby-blog/internal/user/domain"

type Post struct {
	ID         uint
	UserID     userDomain.ID
	CategoryID uint
	Category   Category
	Title      string
	Content    string
	Status     Status
}

type Category struct {
	ID   uint
	Name string
}
