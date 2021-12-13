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
	accCipher entity.AccountCipher,
	userRepo repository.UserRepository,
	postRepo repository.PostRepository, saveDir string,
	reportRepo repository.ReportRepository,
	likeRepo repository.LikeRepository,
	relationRepo repository.RelationRepository,
	commentRepo repository.CommentRepository,
) Application {

	return Application{
		UserHandler: NewUserHandler(
			entity.NewUserFactory(accCipher),
			userRepo,
			relationRepo,
			accCipher,
		),
		PostHandler: NewPostHandler(
			postRepo,
			saveDir,
			likeRepo,
			commentRepo,
			relationRepo,
			userRepo,
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
