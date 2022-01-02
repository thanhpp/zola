package controller

import (
	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
)

var (
	ErrInvalidMediaURL = errors.New("invalid media URL")
	ErrEmptyMediaURL   = errors.New("empty media URL")
)

type ResolveMediaURLFn func(url string) (postID, mediaID string, err error)

type UserController struct {
	handler            application.UserHandler
	postHdl            application.PostHandler
	authsrv            auth.AuthService
	resolveMediaUrlFn  ResolveMediaURLFn
	formUserMediaUrlFn dto.FormUserMediaFn
}

func NewUserCtrl(
	userHandler application.UserHandler,
	postHandler application.PostHandler,
	authSrv auth.AuthService,
	resolveMediaURLFn ResolveMediaURLFn,
	formUserMediaUrlFn dto.FormUserMediaFn,
) *UserController {
	return &UserController{
		handler:            userHandler,
		postHdl:            postHandler,
		authsrv:            authSrv,
		resolveMediaUrlFn:  resolveMediaURLFn,
		formUserMediaUrlFn: formUserMediaUrlFn,
	}
}
