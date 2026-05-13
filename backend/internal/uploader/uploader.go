package uploader

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"path/filepath"

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
	log.Printf("[DEBUG] File: %s, ContentType: %s, Size: %d bytes", file.Filename, contentType, file.Size)

	var mediaType media.Type
	switch {
	case strings.HasPrefix(contentType, "image/"):
		mediaType = media.TypeImage
	case strings.HasPrefix(contentType, "video/"):
		mediaType = media.TypeVideo
	default:
		log.Printf("[ERROR] Unsupported media type: %s", contentType)
		return "", "", appErrors.ErrUnsupportedMedia
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return "", "", err
	}

	ext := filepath.Ext(file.Filename)
	if ext == "" {
		log.Printf("[ERROR] No file extension: %s", file.Filename)
		return "", "", appErrors.ErrUnsupportedMedia
	}

	filename := uuid.New().String() + ext
	log.Printf("[DEBUG] Generated filename: %s", filename)

	path, err := u.storage.Save(f, filename)
	if err != nil {
		return "", "", err
	}

	return path, mediaType, nil
}

func (u *Uploader) Delete(path string) error {
	return u.storage.Delete(path)
}
