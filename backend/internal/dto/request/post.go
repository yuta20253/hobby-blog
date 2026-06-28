package request

import (
	"hobby-blog/internal/model"
	"hobby-blog/internal/service/input"
)

type PostSearchRequest struct {
	Title    string `form:"title"`
	UserName string `form:"user_name"`
	Category string `form:"category"`
	Limit    int    `form:"limit"`
	Offset   int    `form:"offset"`
}

type CreatePostRequest struct {
	Title      string `json:"title" binding:"required,max=255"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

type UpdatePostRequest struct {
	Title      string       `json:"title" binding:"required,max=255"`
	Content    string       `json:"content" binding:"required"`
	CategoryID uint         `json:"category_id" binding:"required"`
	Status     domain.Status `json:"status" binding:"required"`
}

func (q PostSearchRequest) ToInput() input.SearchPostQuery {
	return input.SearchPostQuery{
		Title:    q.Title,
		UserName: q.UserName,
		Category: q.Category,
		Limit:    q.Limit,
		Offset:   q.Offset,
	}
}

func (r CreatePostRequest) ToInput(userID uint) input.CreatePostInput {
	return input.CreatePostInput{
		Title:      r.Title,
		Content:    r.Content,
		CategoryID: r.CategoryID,
		UserID:     userID,
	}
}

func (r UpdatePostRequest) ToInput(id uint, userID uint) input.UpdatePostInput {
	return input.UpdatePostInput{
		ID:         id,
		Title:      r.Title,
		Content:    r.Content,
		CategoryID: r.CategoryID,
		UserID:     userID,
		Status:     r.Status,
	}
}
