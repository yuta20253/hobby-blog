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
	uid, ok := getUserID(c)

	if !ok {
		return
	}

	mypage, err := h.service.GetMyPage(uid)

	if err != nil {
		respondError(c, 404, "mypage not found")
		return
	}

	c.JSON(200, gin.H{
		"mypage": mypage,
	})
}
