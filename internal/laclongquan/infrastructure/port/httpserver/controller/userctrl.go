package controller

import (
	"github.com/thanhpp/zola/internal/laclongquan/application"
)

type UserController struct {
	handler application.UserHandler
}

func NewUserCtrl(userHandler application.UserHandler) *UserController {
	return &UserController{
		handler: userHandler,
	}
}
