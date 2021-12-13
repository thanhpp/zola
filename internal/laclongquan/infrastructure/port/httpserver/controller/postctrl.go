package controller

import (
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
)

type PostController struct {
	handler          application.PostHandler
	likeHandler      application.LikeHandler
	formMediaURLFunc dto.FormMediaURLFunc
}

func NewPostCtrl(h application.PostHandler, likeHdl application.LikeHandler, formURLFn dto.FormMediaURLFunc) *PostController {
	return &PostController{
		handler:          h,
		likeHandler:      likeHdl,
		formMediaURLFunc: formURLFn,
	}
}
