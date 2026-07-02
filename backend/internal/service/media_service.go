package service

import (
	"context"
	stdErrors "errors"
	"mime/multipart"
	"path/filepath"

	"gorm.io/gorm"
	appErrors "hobby-blog/internal/errors"
	postDomain "hobby-blog/internal/post/domain"
	"hobby-blog/internal/repository"
	"hobby-blog/internal/uploader"
	mediaModel "hobby-blog/internal/model"
)

type MediaService struct {
	postRepo  postDomain.PostRepository
	mediaRepo repository.MediaRepository
	uploader  *uploader.Uploader
}

func NewMediaService(
	postRepo postDomain.PostRepository,
	mediaRepo repository.MediaRepository,
	uploader *uploader.Uploader,
) *MediaService {
	return &MediaService{
		postRepo:  postRepo,
		mediaRepo: mediaRepo,
		uploader:  uploader,
	}
}

func (s *MediaService) CreateMedia(ctx context.Context, userID uint, postID uint, file *multipart.FileHeader) error {
	post, err := s.postRepo.GetByID(ctx, postID)
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return appErrors.ErrNotFound
		}
		return err
	}

	if uint(post.UserID) != userID {
		return appErrors.ErrForbidden
	}

	filePath, mediaType, err := s.uploader.Upload(file)
	if err != nil {
		return err
	}

	err = s.mediaRepo.Create(&mediaModel.MediaFile{
		PostID:   postID,
		Type:     mediaType,
		FilePath: filePath,
		FileName: filepath.Base(filePath),
	})

	if err != nil {
		_ = s.uploader.Delete(filePath)
		return err
	}

	return nil
}
