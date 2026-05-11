package service

import (
	"errors"
	"hobby-blog/internal/domain/media"
	"hobby-blog/internal/model"
	"hobby-blog/internal/repository"
)

type MediaService struct {
	postRepo  *repository.PostRepository
	mediaRepo *repository.MediaRepository
}

var ErrForbidden = errors.New("forbidden")

func NewMediaService(
	postRepo *repository.PostRepository,
	mediaRepo *repository.MediaRepository,
) *MediaService {
	return &MediaService{
		postRepo:  postRepo,
		mediaRepo: mediaRepo,
	}
}

func (s *MediaService) CreateMedia(userID uint, postID int, path string, fileName string, mediaType media.Type) error {
	post, err := s.postRepo.FindByID(postID)

	if err != nil {
		return err
	}

	if post.UserID != userID {
		return ErrForbidden
	}

	return s.mediaRepo.Create(&model.MediaFile{
		PostID:   uint(postID),
		Type:     mediaType,
		FilePath: path,
		FileName: fileName,
	})
}
