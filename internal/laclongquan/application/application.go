package application

import (
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type Application struct {
	UserHandler   UserHandler
	PostHandler   PostHandler
	ReportHandler ReportHandler
	LikeHandler   LikeHandler
}

func NewApplication(
	userRepo repository.UserRepository,
	postRepo repository.PostRepository, saveDir string,
	reportRepo repository.ReportRepository,
	likeRepo repository.LikeRepository,
) Application {

	return Application{
		UserHandler: NewUserHandler(
			entity.NewUserFactory(),
			userRepo,
		),
		PostHandler: NewPostHandler(
			postRepo,
			saveDir,
		),
		ReportHandler: NewReportHandler(
			entity.NewReportFactory(),
			reportRepo,
			postRepo,
		),
		LikeHandler: NewLikeHandler(
			entity.NewLikeFactory(),
			likeRepo,
			postRepo,
		),
	}
}
