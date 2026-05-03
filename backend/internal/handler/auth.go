package handler

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) SignUp(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "signup ok",
	})
}
