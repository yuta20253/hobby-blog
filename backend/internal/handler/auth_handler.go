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

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
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

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	result, err := h.service.Login(req.Email, req.Password)
	if (err != nil) {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(200, gin.H{
		"user": result.User,
		"token": result.Token,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout success",
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	id, ok := userID.(uint)

	if !ok {
		c.JSON(500, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "user not found"
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
