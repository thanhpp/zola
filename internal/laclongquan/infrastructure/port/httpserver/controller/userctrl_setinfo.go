package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

const (
	avatarKey     = "avatar"
	coverImageKey = "cover_image"
)

func (ctrl UserController) SetUserInfo(c *gin.Context) {
	var req = new(dto.SetUserInfoReq)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind request error: %v", err)
		ginAbortInternalError(c, responsevalue.CodeInvalidParameterType, responsevalue.MsgInvalidRequest, req)
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("Error while getting userID: %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	var (
		avatarMedia *entity.Media
		coverMedia  *entity.Media
	)

	avaPostID, avaMediaID, err := ctrl.resolveMediaUrlFn(req.Avatar)
	if err != nil && !errors.Is(err, ErrEmptyMediaURL) {
		logger.Errorf("resolve media url (avatar) error: %v", err)
		ginAbortInternalError(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
		return
	}
	if len(avaPostID)+len(avaMediaID) != 0 {
		avatarMedia, err = ctrl.postHdl.GetMedia(c, userID.String(), avaPostID, avaMediaID)
		if err != nil {
			logger.Errorf("get media (avatar) error: %v", err)
			ginAbortInternalError(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
			return
		}
	}

	coverPostID, coverMediaID, err := ctrl.resolveMediaUrlFn(req.CoverImage)
	if err != nil && !errors.Is(err, ErrEmptyMediaURL) {
		logger.Errorf("resolve media url (cover image) error: %v", err)
		ginAbortInternalError(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
		return
	}
	if len(coverPostID)+len(coverMediaID) != 0 {
		coverMedia, err = ctrl.postHdl.GetMedia(c, userID.String(), coverPostID, coverMediaID)
		if err != nil {
			logger.Errorf("get media (cover image) error: %v", err)
			ginAbortInternalError(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
			return
		}
	}

	err = ctrl.handler.SetUserInfo(
		c, userID,
		req.Username, req.Description,
		req.Address, req.City, req.Country,
		req.Link,
		avatarMedia, coverMedia,
	)
	if err != nil {
		logger.Errorf("set user info %s: %v", userID.String(), err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
			return

		case entity.ErrInvalidInputLength:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
			return

		case entity.ErrInvalidUsername:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
			return

		case entity.ErrInvalidCountry:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
			return

		case application.ErrCanNotUseMedia:
			ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, responsevalue.MsgInvalidRequest, nil)
			return
		}
		ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, req)
		return
	}

	var resp = new(dto.SetUserInfoResp)
	resp.SetCode(responsevalue.CodeOK)
	resp.SetMsg(responsevalue.MsgOK)
	c.JSON(http.StatusOK, resp)
}
