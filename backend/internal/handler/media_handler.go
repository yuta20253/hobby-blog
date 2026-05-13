package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/service"
	"strconv"
)

type MediaHandler struct {
	service *service.MediaService
}

func NewMediaHandler(service *service.MediaService) *MediaHandler {
	return &MediaHandler{
		service: service,
	}
}

func (h *MediaHandler) UploadMedia(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	uid, ok := userID.(uint)

	if !ok {
		c.JSON(500, gin.H{"error": "invalid user id"})
		return
	}

	strPostID := c.Param("id")

	postID, err := strconv.Atoi(strPostID)

	if err != nil {
		c.JSON(500, gin.H{"error": "invalid id"})
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}

	err = h.service.CreateMedia(uid, uint(postID), file)

	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrForbidden):
			c.JSON(403, gin.H{"error": "forbidden"})
		case errors.Is(err, appErrors.ErrUnsupportedMedia):
			c.JSON(400, gin.H{"error": "unsupported file type"})
		case errors.Is(err, appErrors.ErrNotFound):
			c.JSON(404, gin.H{"error": "post not found"})
		default:
			c.JSON(500, gin.H{"error": "failed to save media"})
		}
		return
	}

	c.JSON(201, gin.H{
		"message": "uploaded",
	})
}
