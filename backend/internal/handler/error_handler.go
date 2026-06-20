package handler

import (
	"errors"
	"net/http"

	appErrors "hobby-blog/internal/errors"
	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error)  {
	var appErr *appErrors.AppError

	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, gin.H{
			"error": appErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
}
