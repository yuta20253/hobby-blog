package service

import (
	stdErrors "errors"
	"mime/multipart"
	"path/filepath"

	"gorm.io/gorm"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/model"
	"hobby-blog/internal/repository"
	"hobby-blog/internal/uploader"
)

type MediaService struct {
	postRepo  repository.PostRepository
	mediaRepo repository.MediaRepository
	uploader  *uploader.Uploader
}

func NewMediaService(
	postRepo repository.PostRepository,
	mediaRepo repository.MediaRepository,
	uploader *uploader.Uploader,
) *MediaService {
	return &MediaService{
		postRepo:  postRepo,
		mediaRepo: mediaRepo,
		uploader:  uploader,
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

	filePath, mediaType, err := s.uploader.Upload(file)
	if err != nil {
		return err
	}

	return s.mediaRepo.Create(&model.MediaFile{
		PostID:   postID,
		Type:     mediaType,
		FilePath: filePath,
		FileName: filepath.Base(filePath),
	})
}
