package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"hobby-blog/internal/handler"
)

func SetUpRouter(h *handler.AuthHandler) *gin.Engine  {
	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	auth := api.Group("/auth")

	auth.POST("/signup", h.SignUp)

	return r
}
