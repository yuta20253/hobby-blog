package response

import "hobby-blog/internal/model"

func NewMediaFileResponse(f model.MediaFile) MediaFileResponse {
	return MediaFileResponse{
		ID:       f.ID,
		Type:     f.Type,
		FilePath: f.FilePath,
		FileName: f.FileName,
	}
}

func NewMediaFileResponses(files []model.MediaFile) []MediaFileResponse {
	res := make([]MediaFileResponse, 0, len(files))

	for _, f := range files {
		res = append(res, NewMediaFileResponse(f))
	}

	return res
}
