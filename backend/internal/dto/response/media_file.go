package response

import "hobby-blog/internal/domain/media"

type MediaFileResponse struct {
	ID       uint       `json:"id"`
	Type     media.Type `json:"type"`
	FilePath string     `json:"file_path"`
	FileName string     `json:"file_name"`
}
