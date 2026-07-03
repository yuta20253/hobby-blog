package presentation

import (
	"github.com/gin-gonic/gin"

	httphelper "hobby-blog/internal/common/http"
	mypageApplication "hobby-blog/internal/mypage/application"
)

type MypageHandler struct {
	service *mypageApplication.MypageService
}

func NewMypageHandler(service *mypageApplication.MypageService) *MypageHandler {
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
