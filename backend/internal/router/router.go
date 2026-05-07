package router

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"hobby-blog/internal/handler"
	"hobby-blog/internal/middleware"
)

func SetUpRouter(
		authHandler *handler.AuthHandler,
		postHandler *handler.PostHandler,
	) *gin.Engine  {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},

		ExposeHeaders: []string{
			"Content-Length",
		},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	if gin.Mode() == gin.DebugMode {
		for _, route := range r.Routes() {
			log.Println(route.Method, route.Path)
		}
	}

	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/signup", authHandler.SignUp)
	auth.POST("/login", authHandler.Login)

	posts := api.Group("/posts")
	posts.GET("", postHandler.Index)
	posts.GET("/:id", postHandler.Show)

	private := api.Group("")
	private.Use(middleware.AuthMiddleware())

	private.DELETE("/logout", authHandler.Logout)
	private.GET("/me", authHandler.Me)

	postsPrivate := private.Group("/posts")
	postsPrivate.POST("", postHandler.Create)
	postsPrivate.PATCH("/:id", postHandler.Update)
	postsPrivate.DELETE("/:id", postHandler.Delete)

	return r
}
