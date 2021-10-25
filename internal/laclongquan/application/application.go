package application

import (
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type Application struct {
	UserHandler UserHandler
}

func NewApplication(userRepo repository.UserRepository) Application {

	return Application{
		UserHandler: NewUserHandler(
			entity.NewUserFactory(),
			userRepo,
		),
	}
}
