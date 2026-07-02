package application

type PostDTO struct {
	ID      uint
	Title   string
	Content string
	Status  string

	User UserDTO

	Category CategoryDTO

	MediaFiles []MediaFileDTO
}

type UserDTO struct {
	ID    uint
	Name  string
	Email string
}

type CategoryDTO struct {
	ID   uint
	Name string
}

type MediaFileDTO struct {
	ID       uint
	Type     string
	FilePath string
	FileName string
}
