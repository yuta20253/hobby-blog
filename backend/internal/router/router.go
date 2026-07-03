package router

import (
	"log"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"hobby-blog/internal/config"
	"hobby-blog/internal/middleware"

	mediaPresentation "hobby-blog/internal/media/presentation"
	mypagePresentation "hobby-blog/internal/mypage/presentation"
	postPresentation "hobby-blog/internal/post/presentation"
	authPresentation "hobby-blog/internal/user/presentation"
)

func SetUpRouter(
	cfg *config.Config,

	authHandler *authPresentation.AuthHandler,
	postHandler *postPresentation.PostHandler,

	mypageHandler *mypagePresentation.MypageHandler,
	mediaHandler *mediaPresentation.MediaHandler,
) *gin.Engine {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSAllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
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

	postsPrivate.POST("/:id/media", mediaHandler.UploadMedia)

	mypage := private.Group("")
	mypage.GET("/mypage", mypageHandler.Show)

	uploadPath, err := filepath.Abs("./uploads")
	if err != nil {
		log.Fatal(err)
	}
	r.Static("/uploads", uploadPath)

	return r
}
