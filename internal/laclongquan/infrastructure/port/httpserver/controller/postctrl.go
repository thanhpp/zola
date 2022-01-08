package controller

import (
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
)

type PostController struct {
	handler               application.PostHandler
	likeHandler           application.LikeHandler
	formMediaURLFunc      dto.FormMediaURLFunc
	formVideoThumbURLFunc dto.FormVideoThumbURLFunc
	formUserMediaURLFunc  dto.FormUserMediaFn
}

func NewPostCtrl(
	h application.PostHandler,
	likeHdl application.LikeHandler,
	formURLFn dto.FormMediaURLFunc,
	formVideoThumbURLFn dto.FormVideoThumbURLFunc,
	formUserMediaURLFunc dto.FormUserMediaFn) *PostController {
	return &PostController{
		handler:               h,
		likeHandler:           likeHdl,
		formMediaURLFunc:      formURLFn,
		formVideoThumbURLFunc: formVideoThumbURLFn,
		formUserMediaURLFunc:  formUserMediaURLFunc,
	}
}
