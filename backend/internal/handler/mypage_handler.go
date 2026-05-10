package handler

import (
	"github.com/gin-gonic/gin"
	"hobby-blog/internal/service"
)

type MypageHandler struct {
	service *service.MypageService
}

func NewMypageHandler(service *service.MypageService) *MypageHandler {
	return &MypageHandler{
		service: service,
	}
}

func (h *MypageHandler) Show(c *gin.Context) {
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

	mypage, err := h.service.GetMyPage(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "mypage not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"mypage": mypage,
	})
}
