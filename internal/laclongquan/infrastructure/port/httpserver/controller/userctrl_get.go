package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) GetUserInfo(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		ginAbortUnauthorized(c, responsevalue.ValueInvalidToken, "invalid token")
		return
	}

	requestedID, err := getUserUUIDFromParam(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid user ID")
	}

	res, err := ctrl.handler.GetUserByID(c, requestorID.String(), requestedID.String())
	if err != nil {
		logger.Errorf("can not get user %s info: %v", requestedID.String(), err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not found")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not found")
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not found")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	var resp = new(dto.GetUserResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(res.User, res.FriendCount, res.IsFriend, ctrl.formUserMediaUrlFn)

	c.JSON(http.StatusOK, resp)
}

func (ctrl UserController) GetUserMedia(c *gin.Context) {
	// requestorID, err := getUserUUIDFromClaims(c)
	// if err != nil {
	// 	ginAbortUnauthorized(c, responsevalue.ValueInvalidToken, "invalid token")
	// 	return
	// }

	requestedID, err := getUserUUIDFromParam(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid user ID")
		return
	}

	mediaID, err := getMediaID(c)
	if err != nil {
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid media ID")
		return
	}

	media, err := ctrl.handler.GetUserMedia(c, "requestorID.String()", requestedID.String(), mediaID)
	if err != nil {
		logger.Errorf("can not get user %s media %s: %v", " requestorID.String()", mediaID, err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "media not found")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "media not found")
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "media not found")
			return

		case repository.ErrMediaNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "media not found")
			return

		case application.ErrCanNotGetUserMedia:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "media not found")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	c.File(media.Path())
}
