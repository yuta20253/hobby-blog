package service

import (
	stdErrors "errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"gorm.io/gorm"
	"hobby-blog/internal/domain/media"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/model"
	"hobby-blog/internal/repository"
	"hobby-blog/internal/storage"
)

type MediaService struct {
	postRepo  repository.PostRepository
	mediaRepo repository.MediaRepository
	storage   storage.FileStorage
}

func NewMediaService(
	postRepo repository.PostRepository,
	mediaRepo repository.MediaRepository,
	storage storage.FileStorage,
) *MediaService {
	return &MediaService{
		postRepo:  postRepo,
		mediaRepo: mediaRepo,
		storage:   storage,
	}
}

func (s *MediaService) CreateMedia(userID uint, postID uint, file *multipart.FileHeader) error {
	post, err := s.postRepo.FindByID(postID)
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return appErrors.ErrNotFound
		}
		return err
	}

	if post.UserID != userID {
		return appErrors.ErrForbidden
	}

	tmp, err := file.Open()

	if err != nil {
		return err
	}

	defer tmp.Close()

	buffer := make([]byte, 512)
	n, err := tmp.Read(buffer)

	if err != nil && !stdErrors.Is(err, io.EOF) {
		return err
	}

	contentType := http.DetectContentType(buffer[:n])

	var mediaType media.Type
	switch {
	case strings.HasPrefix(contentType, "image/"):
		mediaType = media.TypeImage
	case strings.HasPrefix(contentType, "video/"):
		mediaType = media.TypeVideo
	default:
		return appErrors.ErrUnsupportedMedia
	}

	path, fileName, err := s.storage.Save(file)

	if err != nil {
		_ = s.storage.Delete(path)
		return err
	}

	return s.mediaRepo.Create(&model.MediaFile{
		PostID:   uint(postID),
		Type:     mediaType,
		FilePath: path,
		FileName: fileName,
	})
}
