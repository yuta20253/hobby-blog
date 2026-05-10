package service

import (
	"hobby-blog/internal/dto"
	"hobby-blog/internal/repository"
)

type MypageService struct {
	userRepo *repository.UserRepository
	postRepo *repository.PostRepository
}

type MypageResponse struct {
	User  UserResponse       `json:"user"`
	Posts []dto.PostResponse `json:"posts"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewMypageService(userRepo *repository.UserRepository, postRepo *repository.PostRepository) *MypageService {
	return &MypageService{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (s *MypageService) GetMyPage(id uint) (*MypageResponse, error) {
	user, err := s.userRepo.FindByID(id)

	if err != nil {
		return nil, err
	}

	posts, err := s.postRepo.GetMyPostsByUserID(user.ID)

	if err != nil {
		return nil, err
	}

	postResponses := make([]dto.PostResponse, 0, len(posts))

	for _, p := range posts {
		postResponses = append(postResponses, dto.PostResponse{
			ID:      p.ID,
			Title:   p.Title,
			Content: p.Content,
			Category: dto.CategoryResponse{
				ID:   p.Category.ID,
				Name: p.Category.Name,
			},
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
