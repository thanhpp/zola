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
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, responsevalue.MsgInvalidRequest)
		return
	}

	blockerID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user uuid from claims error %v", err)
		ginAbortInternalError(c, responsevalue.ValueInvalidateUser, "invalid user")
		return
	}

	switch req.Type {
	case BlockCode:
		if err := ctrl.handler.BlockDiary(c, blockerID.String(), req.BlockedUserID); err != nil {
			logger.Errorf("block diary issuer %s issued %s error %v", blockerID.String(), req.BlockedUserID, err)
			switch err {
			case repository.ErrUserNotFound:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not exist")
				return

			case repository.ErrRelationNotFound:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "relation not exist")
				return

			case entity.ErrSelfRelation:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "can't block yourself")
				return

			case entity.ErrLockedUser:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user has been locked")
				return

			case entity.ErrAlreadyBlocked:
				ginAbortNotAcceptable(c, responsevalue.ValueActionHasBeenDone, "user has already been blocked")
				return

			case repository.ErrSameUser:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "same user")
				return
			}
			ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
			return
		}
		ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
		return

	case UnblockCode:
		if err := ctrl.handler.UnblockDiary(c, blockerID.String(), req.BlockedUserID); err != nil {
			logger.Errorf("unblock diary issuer %s issued %s error %v", blockerID.String(), req.BlockedUserID, err)
			switch err {
			case repository.ErrUserNotFound:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not exist")
				return

			case repository.ErrRelationNotFound:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "relation not exist")
				return

			case entity.ErrLockedUser:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user has been locked")
				return

			case entity.ErrInvalidRelation:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "can not unblock diary")
				return

			case repository.ErrSameUser:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "same user")
				return
			}
			ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
			return
		}
		ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
		return

	default:
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid block type")
		return
	}
}
