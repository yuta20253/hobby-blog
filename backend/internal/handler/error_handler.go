package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"hobby-blog/internal/dto/response"
	appErrors "hobby-blog/internal/errors"
)

func handleError(c *gin.Context, err error) {
	var appErr *appErrors.AppError

	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, response.ErrorResponse{
			Error: appErr.Message,
		})
		return
	}

	c.Error(err)

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
}
