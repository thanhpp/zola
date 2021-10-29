package controller

import (
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
)

type UserController struct {
	handler application.UserHandler
	authsrv auth.AuthService
}

func NewUserCtrl(userHandler application.UserHandler, authSrv auth.AuthService) *UserController {
	return &UserController{
		handler: userHandler,
		authsrv: authSrv,
	}
}
