package presentation

import (
	"github.com/gin-gonic/gin"
	httphelper "hobby-blog/internal/common/http"
	userApplicationUsecase "hobby-blog/internal/user/application"
)

type AuthHandler struct {
	service *userApplicationUsecase.AuthService
}

func NewAuthHandler(service *userApplicationUsecase.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.SignUp(
		c.Request.Context(),
		userApplicationUsecase.SignUpInput{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		},
	)

	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	c.JSON(200, AuthResponse{
		User:  NewAuthUserResponse(result.User),
		Token: result.Token,
	})

}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		httphelper.HandleError(c, err)
		return
	}

	result, err := h.service.Login(c.Request.Context(), userApplicationUsecase.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, AuthResponse{
		User:  NewAuthUserResponse(result.User),
		Token: result.Token,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout success",
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	uid, ok := httphelper.GetUserID(c)

	if !ok {
		return
	}

	user, err := h.service.GetUserByID(c.Request.Context(), uid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user": NewAuthUserResponse(user),
	})
}
