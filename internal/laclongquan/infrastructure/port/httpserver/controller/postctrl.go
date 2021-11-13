package controller

import (
	"github.com/thanhpp/zola/internal/laclongquan/application"
)

type PostController struct {
	handler     application.PostHandler
	likeHandler application.LikeHandler
}

func NewPostCtrl(h application.PostHandler, likeHdl application.LikeHandler) *PostController {
	return &PostController{
		handler:     h,
		likeHandler: likeHdl,
	}
}
