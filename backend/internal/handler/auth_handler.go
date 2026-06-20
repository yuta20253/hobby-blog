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
		respondError(c, 400, "invalid request")
		return
	}

	result, err := h.service.SignUp(req.Name, req.Email, req.Password)

	if err != nil {
		respondError(c, 500, "failed")
		return
	}

	c.JSON(200, gin.H{
		"user":  result.User,
		"token": result.Token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, 400, "invalid request")
		return
	}

	result, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"user":  result.User,
		"token": result.Token,
	})
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
