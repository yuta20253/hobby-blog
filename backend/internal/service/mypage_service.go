package service

import (
	"errors"
	"gorm.io/gorm"
	"hobby-blog/internal/dto/response"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/repository"
)

type MypageService struct {
	userRepo repository.UserRepository
	postRepo repository.PostRepository
}

func NewMypageService(userRepo repository.UserRepository, postRepo repository.PostRepository) *MypageService {
	return &MypageService{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (s *MypageService) GetMyPage(id uint) (*response.MypageResponse, error) {
	user, err := s.userRepo.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}

		return nil, err
	}

	posts, err := s.postRepo.GetMyPostsByUserID(user.ID)

	if err != nil {
		return nil, err
	}

	return response.NewMypageResponse(user, posts), nil
}
