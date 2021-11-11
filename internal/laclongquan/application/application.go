package application

import (
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type Application struct {
	UserHandler   UserHandler
	PostHandler   PostHandler
	ReportHandler ReportHandler
}

func NewApplication(
	userRepo repository.UserRepository,
	postRepo repository.PostRepository, saveDir string,
	reportRepo repository.ReportRepository,
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
	}
}
