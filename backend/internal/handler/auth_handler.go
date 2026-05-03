package handler

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"hobby-blog/internal/service"
)


type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		service: service.NewAuthService(db),
	}
}

func (h *AuthHandler) SignUp(c *gin.Context)  {
	err := h.service.SignUp()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "signup ok",
	})
}
