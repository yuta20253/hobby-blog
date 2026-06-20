package handler

import (
	"github.com/gin-gonic/gin"

	"hobby-blog/internal/dto/request"
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
		handleError(c, err)
		return
	}

	posts, err := h.service.SearchPosts(q.ToInput())

	if err != nil {
		handleError(c, err)
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
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func (h *PostHandler) Create(c *gin.Context) {
	var req request.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	uid, ok := getUserID(c)

	if !ok {
		return
	}

	err := h.service.CreatePost(req.ToInput(uid))

	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "success create post",
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	var req request.UpdatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	uid, ok := getUserID(c)

	if !ok {
		return
	}

	postID, ok := getParamID(c, "id")

	if !ok {
		return
	}

	updatedPost, err := h.service.UpdatePost(req.ToInput(postID, uid))

	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"post": updatedPost,
	})
}

func (h *PostHandler) Delete(c *gin.Context) {
	uid, ok := getUserID(c)

	if !ok {
		return
	}

	postID, ok := getParamID(c, "id")

	if !ok {
		return
	}

	err := h.service.DeletePost(postID, uid)

	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "success delete post"})
}
