package httphelper

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"hobby-blog/internal/dto/response"
	appErrors "hobby-blog/internal/errors"
)

func HandleError(c *gin.Context, err error) {
	var appErr *appErrors.AppError

	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, response.ErrorResponse{
			Error: appErr.Message,
		})
		return
	}

	var validationErr validator.ValidationErrors

	if errors.As(err, &validationErr) {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error: "invalid request",
		})
		return
	}

	log.Println(err)

	c.JSON(http.StatusInternalServerError, response.ErrorResponse{
		Error: "internal server error",
	})
}
