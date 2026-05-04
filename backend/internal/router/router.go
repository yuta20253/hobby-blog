package router

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"hobby-blog/internal/handler"
)

func SetUpRouter(h *handler.AuthHandler) *gin.Engine  {
	r := gin.Default()
	r.Use(cors.Default())

	if gin.Mode() == gin.DebugMode {
		for _, route := range r.Routes() {
			log.Println(route.Method, route.Path)
		}
	}

	api := r.Group("/api")
	auth := api.Group("/auth")

	auth.POST("/signup", h.SignUp)
	auth.POST("/login", h.Login)

	authWithAuth := auth.Group("")
	authWithAuth.Use(middleware.AuthMiddleware())

	auth.DELETE("/logout", h.Logout)
	auth.GET("/me", h.Me)

	return r
}
