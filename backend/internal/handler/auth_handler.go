package handler

import (
	"github.com/gin-gonic/gin"
	"hobby-blog/internal/dto/request"
	"hobby-blog/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req request.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	result, err := h.service.SignUp(req.ToInput())

	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, result)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	result, err := h.service.Login(req.ToInput())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, result)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout success",
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	uid, ok := getUserID(c)

	if !ok {
		return
	}

	user, err := h.service.GetUserByID(uid)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
