package domain

type Type string

const (
	TypeImage Type = "image"
	TypeVideo Type = "video"
)

type Media struct {
	ID       uint
	PostID   uint
	Type     Type
	FilePath string
	FileName string
}
