package container

import (
	"gorm.io/gorm"

	"hobby-blog/internal/config"

	mediaPresentation "hobby-blog/internal/media/presentation"
	mypagePresentation "hobby-blog/internal/mypage/presentation"
	postPresentation "hobby-blog/internal/post/presentation"
	userPresentation "hobby-blog/internal/user/presentation"

	mediaApplication "hobby-blog/internal/media/application"
	mypageApplication "hobby-blog/internal/mypage/application"
	postApplication "hobby-blog/internal/post/application"
	userApplication "hobby-blog/internal/user/application"

	mediaInfrastructure "hobby-blog/internal/media/infrastructure"
	postInfrastructure "hobby-blog/internal/post/infrastructure"
	userInfrastructure "hobby-blog/internal/user/infrastructure"

	"hobby-blog/internal/storage"
	"hobby-blog/internal/uploader"
)

type Container struct {
	AuthHandler   *userPresentation.AuthHandler
	PostHandler   *postPresentation.PostHandler
	MypageHandler *mypagePresentation.MypageHandler
	MediaHandler  *mediaPresentation.MediaHandler
}

func NewContainer(db *gorm.DB) *Container {
	cfg := config.Load()

	userRepo := userInfrastructure.NewUserRepository(db)
	authService := userApplication.NewAuthService(userRepo)
	authHandler := userPresentation.NewAuthHandler(authService)

	postRepo := postInfrastructure.NewPostRepository(db)
	postService := postApplication.NewPostService(postRepo)
	postHandler := postPresentation.NewPostHandler(postService)

	mypageService := mypageApplication.NewMypageService(userRepo, postRepo)
	mypageHandler := mypagePresentation.NewMypageHandler(mypageService)

	mediaRepo := mediaInfrastructure.NewMediaRepository(db)

	st := storage.NewLocalStorage(cfg.UploadPath)
	upl := uploader.NewUploader(st)

	mediaService := mediaApplication.NewMediaService(postRepo, mediaRepo, upl)
	mediaHandler := mediaPresentation.NewMediaHandler(mediaService)

	return &Container{
		AuthHandler:   authHandler,
		PostHandler:   postHandler,
		MypageHandler: mypageHandler,
		MediaHandler:  mediaHandler,
	}
}
