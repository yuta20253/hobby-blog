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

type MypageResponse struct {
	User  UserResponse            `json:"user"`
	Posts []response.PostResponse `json:"posts"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewMypageService(userRepo repository.UserRepository, postRepo repository.PostRepository) *MypageService {
	return &MypageService{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (s *MypageService) GetMyPage(id uint) (*MypageResponse, error) {
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

	postResponses := make([]response.PostResponse, 0, len(posts))

	for _, p := range posts {
		postResponses = append(postResponses, response.PostResponse{
			ID:      p.ID,
			Title:   p.Title,
			Content: p.Content,
			User: response.PostUserResponse{
				ID:    p.User.ID,
				Name:  p.User.Name,
				Email: p.User.Email,
			},
			Category: response.CategoryResponse{
				ID:   p.Category.ID,
				Name: p.Category.Name,
			},
			MediaFiles: response.NewMediaFileResponses(p.MediaFiles),
		})
	}

	return &MypageResponse{
		User: UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Posts: postResponses,
	}, nil
}
