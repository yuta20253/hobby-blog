package handler

import (
	"github.com/gin-gonic/gin"
	"hobby-blog/internal/service"
	"hobby-blog/internal/domain/post"
	"strconv"
	"errors"
	"gorm.io/gorm"
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
	Title string `json:"title" binding:"required,max=255"`
	Content string `json:"content" binding:"required"`
	CategoryID uint `json:"category_id" binding:"required"`
}

type UpdatePostRequest struct {
	Title      string      `json:"title" binding:"required,max=255"`
	Content    string      `json:"content" binding:"required"`
	CategoryID uint        `json:"category_id" binding:"required"`
	Status     post.Status `json:"status" binding:"required"`
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) Index(c *gin.Context) {
	var q PostSearchQuery

	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid query",
		})
		return
	}

	query := post.SearchQuery{
		Title:    q.Title,
		UserName: q.UserName,
		Category: q.Category,
		Limit:    q.Limit,
		Offset:   q.Offset,
	}

	posts, err := h.service.SearchPosts(query)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to fetch posts",
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func (h *PostHandler) Show(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	post, err := h.service.GetPost(uint(id))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "failed to fetch post",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func (h *PostHandler)Create(c *gin.Context) {
	var req CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	id, ok := userID.(uint)

	if !ok {
		c.JSON(500, gin.H{"error": "invalid user id"})
		return
	}

	input := post.CreateInput{
		Title: req.Title,
		Content: req.Content,
		CategoryID: req.CategoryID,
		UserID: id,
	}

	err := h.service.CreatePost(input)

	if err != nil {
		c.JSON(500, gin.H{ "error": "failed" })
		return
	}

	c.JSON(201, gin.H{
		"message": "success create post",
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	var req UpdatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{ "error": "invalid id" })
		return
	}

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{ "error": "unauthorized" })
		return
	}

	uid, ok := userID.(uint)

	if !ok {
		c.JSON(500, gin.H{ "error": "failed" })
		return
	}

	input := post.UpdateInput{
		ID: uint(id),
		Title: req.Title,
		Content: req.Content,
		CategoryID: req.CategoryID,
		UserID: uid,
		Status: req.Status,
	}

	updatedPost, err := h.service.UpdatePost(input)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}

		c.JSON(500, gin.H{ "error": "failed" })
		return
	}

	c.JSON(200, gin.H{
		"post": updatedPost,
	})
}
