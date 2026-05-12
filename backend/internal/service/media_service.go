package service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
	"hobby-blog/internal/domain/media"
	"hobby-blog/internal/model"
	"hobby-blog/internal/repository"
)

type MediaService struct {
	postRepo  *repository.PostRepository
	mediaRepo *repository.MediaRepository
}

var (
	ErrForbidden        = errors.New("forbidden")
	ErrUnsupportedMedia = errors.New("unsupported media")
	ErrNotFound         = errors.New("not found")
)

func NewMediaService(
	postRepo *repository.PostRepository,
	mediaRepo *repository.MediaRepository,
) *MediaService {
	return &MediaService{
		postRepo:  postRepo,
		mediaRepo: mediaRepo,
	}
}

func (s *MediaService) CreateMedia(userID uint, postID int, file *multipart.FileHeader) error {
	post, err := s.postRepo.FindByID(postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}

	if post.UserID != userID {
		return ErrForbidden
	}

	tmp, err := file.Open()

	if err != nil {
		return err
	}

	defer tmp.Close()

	buffer := make([]byte, 512)
	n, err := tmp.Read(buffer)

	if err != nil && !errors.Is(err, io.EOF) {
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
		return ErrUnsupportedMedia
	}

	src, err := file.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := "uploads/" + fileName

	if err := os.MkdirAll("uploads", 0755); err != nil {
		return err
	}

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return err
	}

	return s.mediaRepo.Create(&model.MediaFile{
		PostID:   uint(postID),
		Type:     mediaType,
		FilePath: path,
		FileName: fileName,
	})
}
