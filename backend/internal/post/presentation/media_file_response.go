package presentation

type MediaFileResponse struct {
	ID       uint   `json:"id"`
	Type     string `json:"type"`
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
}
