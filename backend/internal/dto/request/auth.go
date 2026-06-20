package request

import "hobby-blog/internal/service/input"

type SignUpRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (r SignUpRequest) ToInput() input.SignUpInput {
	return input.SignUpInput{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

func (r LoginRequest) ToInput() input.LoginInput {
	return input.LoginInput{
		Email:    r.Email,
		Password: r.Password,
	}
}
