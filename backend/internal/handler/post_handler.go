package handler

import (
	"github.com/gin-gonic/gin"
	"hobby-blog/internal/service"
)

type PostHandler struct {
	service *service.PostService
}

type PostSearchQuery struct {
	Title    string       `form:"title"`
	UserName string       `form:"user_name"`
	Category string       `form:"category"`
	Limit    int          `form:"limit"`
	Offset   int          `form:"offset"`
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

	serviceQuery := service.PostSearchQuery{
		Title:    q.Title,
		UserName: q.UserName,
		Category: q.Category,
		Limit:    q.Limit,
		Offset:   q.Offset,
	}

	posts, err := h.service.SearchPosts(serviceQuery)

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
