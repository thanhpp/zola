package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

var (
	BlockCode   = "0"
	UnblockCode = "1"
)

func (ctrl UserController) BlockUser(c *gin.Context) {
	var req = new(dto.BlockUserReq)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, responsevalue.MsgInvalidRequest)
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user uuid %v", err)
		ginAbortInternalError(c, responsevalue.ValueInvalidateUser, "invalid user id")
		return
	}

	switch req.Type {
	case BlockCode:
		if err := ctrl.handler.BlockUser(c, userID, req.BlockedUserID); err != nil {
			logger.Errorf("block user %v", err)
			switch err {
			case repository.ErrUserNotFound:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not exist")
				return

			case application.ErrAlreadyBlocked:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user has already been blocked")
				return

			case entity.ErrSelfRelation:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "can't block yourself")
				return

			default:
				ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
				return
			}
		}

		ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
		return

	case UnblockCode:
		if err := ctrl.handler.UnblockUser(c, userID, req.BlockedUserID); err != nil {
			logger.Errorf("unblock user %v", err)
			switch err {
			case repository.ErrUserNotFound:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "user not exist")
				return

			case repository.ErrRelationNotFound:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "no relationship")
				return

			case application.ErrNotABlockRelation:
				ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "not a block relation")
				return

			default:
				ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
				return
			}
		}
		ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
		return

	default:
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterValue, "invalid block type")
		return
	}
}
