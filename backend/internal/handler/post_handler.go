package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hobby-blog/internal/dto/request"
	"hobby-blog/internal/model"
	"hobby-blog/internal/service"
)

type PostHandler struct {
	service *service.PostService
}

type PostSearchQuery struct {
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
	Status     model.Status `json:"status" binding:"required"`
}

func (q PostSearchQuery) ToRequest() request.SearchPostQuery {
	return request.SearchPostQuery{
		Title:    q.Title,
		UserName: q.UserName,
		Category: q.Category,
		Limit:    q.Limit,
		Offset:   q.Offset,
	}
}

func (r CreatePostRequest) ToRequest(userID uint) request.CreatePostInput {
	return request.CreatePostInput{
		Title:      r.Title,
		Content:    r.Content,
		CategoryID: r.CategoryID,
		UserID:     userID,
	}
}

func (r UpdatePostRequest) ToRequest(id uint, userID uint) request.UpdatePostInput {
	return request.UpdatePostInput{
		ID:         id,
		Title:      r.Title,
		Content:    r.Content,
		CategoryID: r.CategoryID,
		UserID:     userID,
		Status:     r.Status,
	}
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) Index(c *gin.Context) {
	var q PostSearchQuery

	if err := c.ShouldBindQuery(&q); err != nil {
		respondError(c, 400, "invalid query")
		return
	}

	posts, err := h.service.SearchPosts(q.ToRequest())

	if err != nil {
		respondError(c, 500, "failed to fetch posts")
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func (h *PostHandler) Show(c *gin.Context) {

	postId, ok := getParamID(c, "id")

	if !ok {
		return
	}

	post, err := h.service.GetPost(postId)

	if err != nil {
		respondError(c, 404, "failed to fetch post")
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func (h *PostHandler) Create(c *gin.Context) {
	var req CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, 400, "invalid request")
		return
	}

	uid, ok := getUserID(c)

	if !ok {
		return
	}

	err := h.service.CreatePost(req.ToRequest(uid))

	if err != nil {
		respondError(c, 500, "failed")
		return
	}

	c.JSON(201, gin.H{
		"message": "success create post",
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	var req UpdatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, 400, "invalid request")
		return
	}

	uid, ok := getUserID(c)

	if !ok {
		c.JSON(500, gin.H{"error": "failed"})
		return
	}

	postID, ok := getParamID(c, "id")

	if !ok {
		return
	}

	updatedPost, err := h.service.UpdatePost(req.ToRequest(postID, uid))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(c, 404, "not found")
			return
		}

		respondError(c, 500, "failed")
		return
	}

	c.JSON(200, gin.H{
		"post": updatedPost,
	})
}

func (h *PostHandler) Delete(c *gin.Context) {
	uid, ok := getUserID(c)

	if !ok {
		respondError(c, 500, "invalid user id")
		return
	}

	postID, ok := getParamID(c, "id")

	if !ok {
		return
	}

	err := h.service.DeletePost(postID, uid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(c, 404, "not found")
			return
		}
		respondError(c, 500, "invalid id")
		return
	}

	c.JSON(200, gin.H{"message": "success delete post"})
}
