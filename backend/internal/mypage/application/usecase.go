package application

import (
	"context"
	"errors"
	"gorm.io/gorm"
	appErrors "hobby-blog/internal/errors"
	postDomain "hobby-blog/internal/post/domain"
	userDomain "hobby-blog/internal/user/domain"
)

type MyPageResult struct {
	User  userDomain.User
	Posts []postDomain.Post
}

type MypageService struct {
	userRepo userDomain.UserRepository
	postRepo postDomain.PostRepository
}

func NewMypageService(userRepo userDomain.UserRepository, postRepo postDomain.PostRepository) *MypageService {
	return &MypageService{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (s *MypageService) GetMyPage(ctx context.Context, id uint) (*MyPageResult, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
		return nil, err
	}

	posts, err := s.postRepo.GetMyPostsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &MyPageResult{
		User:  *user,
		Posts: posts,
	}, nil
}
