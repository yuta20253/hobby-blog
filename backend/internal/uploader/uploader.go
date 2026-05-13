package uploader

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"hobby-blog/internal/domain/media"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/storage"
)

type Uploader struct {
	storage storage.FileStorage
}

func NewUploader(st storage.FileStorage) *Uploader {
	return &Uploader{
		storage: st,
	}
}

func (u *Uploader) Upload(file *multipart.FileHeader) (string, media.Type, error) {
	f, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	buffer := make([]byte, 512)
	n, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return "", "", err
	}

	contentType := http.DetectContentType(buffer[:n])

	var mediaType media.Type
	switch {
	case strings.HasPrefix(contentType, "image/"):
		mediaType = media.TypeImage
	case strings.HasPrefix(contentType, "video/"):
		mediaType = media.TypeVideo
	default:
		return "", "", appErrors.ErrUnsupportedMedia
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return "", "", err
	}

	path, err := u.storage.Save(f, file.Filename)
	if err != nil {
		_ = u.storage.Delete(path)
		return "", "", err
	}

	return path, mediaType, nil
}
