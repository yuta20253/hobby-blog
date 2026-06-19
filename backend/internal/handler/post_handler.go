package handler

import (
	"errors"
	"github.com/gin-gonic/gin"

	"hobby-blog/internal/dto/request"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/service"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) Index(c *gin.Context) {
	var q request.PostSearchRequest

	if err := c.ShouldBindQuery(&q); err != nil {
		respondError(c, 400, "invalid query")
		return
	}

	posts, err := h.service.SearchPosts(q.ToInput())

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
		if errors.Is(err, appErrors.ErrNotFound) {
			respondError(c, 404, "not found")
			return
		}

		respondError(c, 500, "failed")
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func (h *PostHandler) Create(c *gin.Context) {
	var req request.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, 400, "invalid request")
		return
	}

	uid, ok := getUserID(c)

	if !ok {
		return
	}

	err := h.service.CreatePost(req.ToInput(uid))

	if err != nil {
		respondError(c, 500, "failed")
		return
	}

	c.JSON(201, gin.H{
		"message": "success create post",
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	var req request.UpdatePostRequest

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

	updatedPost, err := h.service.UpdatePost(req.ToInput(postID, uid))

	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			respondError(c, 404, "not found")
			return
		}

		if errors.Is(err, appErrors.ErrInvalidInput) {
			respondError(c, 400, "invalid input")
			return
		}

		respondError(c, 500, "invalid user id")
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
		if errors.Is(err, appErrors.ErrNotFound) {
			respondError(c, 404, "not found")
			return
		}
		respondError(c, 500, "failed")
		return
	}

	c.JSON(200, gin.H{"message": "success delete post"})
}
