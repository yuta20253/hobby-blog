package handler

import (
	"github.com/gin-gonic/gin"
	httphelper "hobby-blog/internal/common/http"
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
	uid, ok := httphelper.GetUserID(c)

	if !ok {
		return
	}

	postID, ok := httphelper.GetParamID(c, "id")

	if !ok {
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	err = h.service.CreateMedia(c.Request.Context(), uid, postID, file)

	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	c.JSON(201, gin.H{
		"message": "uploaded",
	})
}
