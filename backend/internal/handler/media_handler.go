package handler

import (
	"github.com/gin-gonic/gin"
	"hobby-blog/internal/service"
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
	uid, ok := getUserID(c)

	if !ok {
		return
	}

	postID, ok := getParamID(c, "id")

	if !ok {
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		handleError(c, err)
		return
	}

	err = h.service.CreateMedia(uid, postID, file)

	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "uploaded",
	})
}
