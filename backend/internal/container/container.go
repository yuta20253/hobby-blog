package container

import (
	"gorm.io/gorm"
	"hobby-blog/internal/config"
	"hobby-blog/internal/handler"
	"hobby-blog/internal/repository"
	postPresentation "hobby-blog/internal/post/presentation"
	userPresentation "hobby-blog/internal/user/presentation"
	postApplication "hobby-blog/internal/post/application"
	userApplicationUsecase "hobby-blog/internal/user/application"
    postInfrastructure "hobby-blog/internal/post/infrastructure"
    userInfrastructure "hobby-blog/internal/user/infrastructure"
	"hobby-blog/internal/storage"
	"hobby-blog/internal/uploader"
)

type Container struct {
	AuthHandler   *userPresentation.AuthHandler
	PostHandler   *postPresentation.PostHandler
	MypageHandler *handler.MypageHandler
	MediaHandler  *handler.MediaHandler
}

func NewContainer(db *gorm.DB) *Container {
	cfg := config.Load()

	userRepo := userInfrastructure.NewUserRepository(db)
	authService := userApplicationUsecase.NewAuthService(userRepo)
	authHandler := userPresentation.NewAuthHandler(authService)

	postRepo := postInfrastructure.NewPostRepository(db)
	postService := postApplication.NewPostService(postRepo)
	postHandler := postPresentation.NewPostHandler(postService)

	mypageService := service.NewMypageService(userRepo, postRepo)
	mypageHandler := handler.NewMypageHandler(mypageService)

	mediaRepo := repository.NewMediaRepository(db)

	st := storage.NewLocalStorage(cfg.UploadPath)
	upl := uploader.NewUploader(st)
	mediaService := service.NewMediaService(postRepo, mediaRepo, upl)
	mediaHandler := handler.NewMediaHandler(mediaService)

	return &Container{
		AuthHandler:   authHandler,
		PostHandler:   postHandler,
		MypageHandler: mypageHandler,
		MediaHandler:  mediaHandler,
	}
}
