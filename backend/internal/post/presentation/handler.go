package presentation

import (
	"github.com/gin-gonic/gin"
	httphelper "hobby-blog/internal/common/http"
	postApplication "hobby-blog/internal/post/application"
)

type PostHandler struct {
	service *postApplication.PostService
}

func NewPostHandler(service *postApplication.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Index(c *gin.Context) {
	var q struct {
		Title    string `form:"title"`
		UserName string `form:"user_name"`
		Category string `form:"category"`
		Limit    int    `form:"limit"`
		Offset   int    `form:"offset"`
	}

	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	posts, err := h.service.SearchPosts(c.Request.Context(), postApplication.SearchPostQuery{
		Title:    q.Title,
		UserName: q.UserName,
		Category: q.Category,
		Limit:    q.Limit,
		Offset:   q.Offset,
	})

	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func (h *PostHandler) Show(c *gin.Context) {

	postId, ok := httphelper.GetParamID(c, "id")

	if !ok {
		return
	}

	post, err := h.service.GetPost(c.Request.Context(), postId)

	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func (h *PostHandler) Create(c *gin.Context) {
	var req CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		httphelper.HandleError(c, err)
		return
	}

	uid, ok := httphelper.GetUserID(c)

	if !ok {
		return
	}

	_, err := h.service.CreatePost(c.Request.Context(), postApplication.CreatePostInput{
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
		UserID:     uid,
	})

	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "success create post",
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	var req UpdatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		httphelper.HandleError(c, err)
		return
	}

	uid, ok := httphelper.GetUserID(c)

	if !ok {
		return
	}

	postID, ok := httphelper.GetParamID(c, "id")

	if !ok {
		return
	}

	updatedPost, err := h.service.UpdatePost(c.Request.Context(), postApplication.UpdatePostInput{
		ID:         postID,
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
		UserID:     uid,
	})

	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"post": updatedPost,
	})
}

func (h *PostHandler) Delete(c *gin.Context) {
	uid, ok := httphelper.GetUserID(c)

	if !ok {
		return
	}

	postID, ok := httphelper.GetParamID(c, "id")

	if !ok {
		return
	}

	err := h.service.DeletePost(c.Request.Context(), postID, uid)

	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "success delete post"})
}
