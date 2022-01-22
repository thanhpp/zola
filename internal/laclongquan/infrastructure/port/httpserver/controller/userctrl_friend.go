package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) NewFriendRequest(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from ctx error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	requesteeID, err := getUserUUIDFromParam(c)
	if err != nil {
		logger.Errorf("get user id from param error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	err = ctrl.handler.NewFriendRequest(c, requestorID.String(), requesteeID.String())
	if err != nil {
		logger.Errorf("new friend request error: %v", err)
		switch err {
		case application.ErrRelationExisted:
			ginAbortNotAcceptable(c, responsevalue.ValueActionHasBeenDone, "relation existed")
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not exist")
			return

		case entity.ErrSelfRelation:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "same id request")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "locked user")
			return

		case repository.ErrSameUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "same user")
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}

func (ctrl UserController) UpdateFriendRequest(c *gin.Context) {
	var req = new(dto.UpdateFriendRequestReq)
	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind request error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid request")
		return
	}

	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from ctx error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	requesteeID, err := getUserUUIDFromParam(c)
	if err != nil {
		logger.Errorf("get user id from param error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	switch {
	case req.IsAcceptCode():
		err = ctrl.handler.UpdateFriendRequest(c, requestorID.String(), requesteeID.String(), true)

	case req.IsRejectCode():
		err = ctrl.handler.UpdateFriendRequest(c, requestorID.String(), requesteeID.String(), false)

	default:
		logger.Errorf("invalid request code: %v", req.IsAccept)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid request code")
		return
	}

	if err != nil {
		switch err {
		case repository.ErrRelationNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "friend request not found")
			return

		case entity.ErrNotAFriendRequest:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "not a friend request")
			return

		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
			return

		case entity.ErrLockedUser:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "locked user")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}

func (ctrl UserController) GetRequestedFriends(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from ctx error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	var requestedID uuid.UUID
	requestedIDStr := strings.ReplaceAll(c.Param("userid"), "/", "")
	// logger.Debugf("get requested friends - requested id str: %s", requestedIDStr)
	if requestedIDStr == "" {
		requestedID = requestorID
	} else {
		requestedID, err = uuid.Parse(requestedIDStr)
		if err != nil {
			logger.Errorf("parse user id from param error: %v", err)
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid user")
			return
		}
	}

	offset, limit := pagination(c)
	results, err := ctrl.handler.GetRequestedFriends(c, requestorID.String(), requestedID.String(), offset, limit)
	if err != nil {
		logger.Errorf("Get requested friends of %s by %s error: %v", requestedID.String(), requestorID.String(), err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not exist")
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "permission denied")
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	var resp = new(dto.GetRequestedFriendsResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(results, ctrl.formUserMediaUrlFn)

	c.JSON(http.StatusOK, resp)
}

func (ctrl UserController) GetFriends(c *gin.Context) {
	requestorID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from ctx error: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	var requestedID uuid.UUID
	requestedIDStr := strings.ReplaceAll(c.Param("userid"), "/", "")
	if requestedIDStr == "" {
		requestedID = requestorID
	} else {
		requestedID, err = uuid.Parse(requestedIDStr)
		if err != nil {
			logger.Errorf("parse user id from param error: %v", err)
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid user")
			return
		}
	}

	offset, limit := pagination(c)
	results, err := ctrl.handler.GetUserFriends(c, requestorID.String(), requestedID.String(), offset, limit)
	if err != nil {
		logger.Errorf("Get friends of %s by %s error: %v", requestedID.String(), requestorID.String(), err)
		switch err {
		case repository.ErrUserNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not exist")
			return

		case entity.ErrPermissionDenied:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, "permission denied")
			return
		}
		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	var resp = new(dto.GetUserFriendsResp)
	resp.SetCode(responsevalue.ValueOK.Code)
	resp.SetMsg(responsevalue.MsgOK)
	resp.SetData(results.Friends, results.Total, ctrl.formUserMediaUrlFn)

	c.JSON(http.StatusOK, resp)
}
