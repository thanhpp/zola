package controller

import (
	"github.com/thanhpp/zola/internal/laclongquan/application"
)

type PostController struct {
	handler application.PostHandler
}

func NewPostCtrl(h application.PostHandler) *PostController {
	return &PostController{
		handler: h,
	}
}
