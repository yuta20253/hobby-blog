package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	appErrors "hobby-blog/internal/errors"
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
		respondError(c, 400, "file required")
		return
	}

	err = h.service.CreateMedia(uid, postID, file)

	if err != nil {
		switch {
		case errors.Is(err, appErrors.ErrForbidden):
			respondError(c, 403, "forbidden")
		case errors.Is(err, appErrors.ErrUnsupportedMedia):
			respondError(c, 400, "unsupported file type")
		case errors.Is(err, appErrors.ErrNotFound):
			respondError(c, 404, "post not found")
		default:
			respondError(c, 500, "failed to save media")
		}
		return
	}

	c.JSON(201, gin.H{
		"message": "uploaded",
	})
}
