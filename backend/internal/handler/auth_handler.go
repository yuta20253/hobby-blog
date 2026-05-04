package handler

import (
	"github.com/gin-gonic/gin"
	"hobby-blog/internal/service"
)


type AuthHandler struct {
	service *service.AuthService
}

type SignUpRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context)  {
	var req SignUpRequest

	if err := c.ShouldBindJSON(&req) ;err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	result, err := h.service.SignUp(req.Name, req.Email, req.Password)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"user":  result.User,
		"token": result.Token,
	})
}
