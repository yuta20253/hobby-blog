package response

import postInfrastructureModel "hobby-blog/internal/post/infrastructure"

func NewMediaFileResponse(f postInfrastructureModel.MediaFile) MediaFileResponse {
	return MediaFileResponse{
		ID:       f.ID,
		Type:     f.Type,
		FilePath: f.FilePath,
		FileName: f.FileName,
	}
}

func NewMediaFileResponses(files []postInfrastructureModel.MediaFile) []MediaFileResponse {
	res := make([]MediaFileResponse, 0, len(files))

	for _, f := range files {
		res = append(res, NewMediaFileResponse(f))
	}

	return res
}
