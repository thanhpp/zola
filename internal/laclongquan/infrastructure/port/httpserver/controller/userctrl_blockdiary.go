package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) BlockDiary(c *gin.Context) {
	var req = new(dto.BlockUserReq)
	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, responsevalue.MsgInvalidRequest, req)
		return
	}

	blockerID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user uuid from claims error %v", err)
		ginAbortInternalError(c, responsevalue.CodeInvalidateUser, "invalid user", nil)
		return
	}

	switch req.Type {
	case BlockCode:
		if err := ctrl.handler.BlockDiary(c, blockerID.String(), req.BlockedUserID); err != nil {
			logger.Errorf("block diary issuer %s issued %s error %v", blockerID.String(), req.BlockedUserID, err)
			switch err {
			case repository.ErrUserNotFound:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user not exist", nil)
				return

			case repository.ErrRelationNotFound:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "relation not exist", nil)
				return

			case entity.ErrSelfRelation:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "can't block yourself", nil)
				return

			case entity.ErrLockedUser:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user has been locked", nil)
				return

			case entity.ErrAlreadyBlocked:
				ginAbortNotAcceptable(c, responsevalue.CodeActionHasBeenDone, "user has already been blocked", nil)
				return

			case repository.ErrSameUser:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "same user", nil)
				return
			}
			ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
			return
		}
		ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
		return

	case UnblockCode:
		if err := ctrl.handler.UnblockDiary(c, blockerID.String(), req.BlockedUserID); err != nil {
			logger.Errorf("unblock diary issuer %s issued %s error %v", blockerID.String(), req.BlockedUserID, err)
			switch err {
			case repository.ErrUserNotFound:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user not exist", nil)
				return

			case repository.ErrRelationNotFound:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "relation not exist", nil)
				return

			case entity.ErrLockedUser:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user has been locked", nil)
				return

			case entity.ErrInvalidRelation:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "can not unblock diary", nil)
				return

			case repository.ErrSameUser:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "same user", nil)
				return
			}
			ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
			return
		}
		ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
		return

	default:
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid block type", nil)
		return
	}
}
