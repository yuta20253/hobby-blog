package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hobby-blog/internal/domain/media"
	"hobby-blog/internal/service"
	"strconv"
	"strings"
	"time"
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

	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	path := "uploads/" + fileName

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(500, gin.H{"error": "failed to save file"})
		return
	}

	contentType := file.Header.Get("Content-Type")

	var mediaType media.Type

	if strings.HasPrefix(contentType, "image/") {
		mediaType = media.TypeImage
	} else if strings.HasPrefix(contentType, "video/") {
		mediaType = media.TypeVideo
	} else {
		c.JSON(400, gin.H{"error": "unsupported file type"})
		return
	}

	err = h.service.CreateMedia(uid, postID, path, fileName, mediaType)

	if err != nil {
		if errors.Is(err, service.ErrForbidden) {
			c.JSON(403, gin.H{"error": "forbidden"})
			return
		}

		c.JSON(500, gin.H{"error": "failed to save media"})
		return
	}

	c.JSON(201, gin.H{
		"message": "uploaded",
	})
}
