package handler

import (
	"github.com/gin-gonic/gin"

	httphelper "hobby-blog/internal/common/http"
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
	uid, ok := httphelper.GetUserID(c)
	if !ok {
		return
	}

	mypage, err := h.service.GetMyPage(c.Request.Context(), uid)
	if err != nil {
		httphelper.HandleError(c, err)
		return
	}

	c.JSON(200, mypage)
}
