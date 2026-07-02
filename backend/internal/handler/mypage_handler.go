package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"

	postPresentation "hobby-blog/internal/post/presentation"
	userPresentation "hobby-blog/internal/user/presentation"
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
		handleError(c, err)
		return
	}

	c.JSON(200, MypageResponse{
		User: userPresentation.NewAuthUserResponse(mypage.User),
		Posts: postPresentation.NewPostResponses(mypage.Posts),
	})
}
