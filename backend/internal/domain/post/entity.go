package domain

type Post struct {
	ID uint
	UserID UserID
	CategoryID uint
	Title string
	Content string
	Status Status
}
