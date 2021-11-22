package controller

import (
	"github.com/gin-gonic/gin"
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
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterType, responsevalue.MsgInvalidRequest, req)
		return
	}

	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user uuid %v", err)
		ginAbortInternalError(c, responsevalue.CodeInvalidateUser, "invalid user id", nil)
		return
	}

	switch req.Type {
	case BlockCode:
		if err := ctrl.handler.BlockUser(c, userID, req.BlockedUserID); err != nil {
			logger.Errorf("block user %v", err)
			switch err {
			case repository.ErrUserNotFound:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user not exist", nil)
				return

			case repository.ErrUserAlreadyBlocked:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user has already been blocked", nil)
				return

			case entity.ErrSelfBlock:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "can't block yourself", nil)
				return

			default:
				ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
				return
			}
		}

		ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
		return

	case UnblockCode:
		if err := ctrl.handler.UnblockUser(c, userID, req.BlockedUserID); err != nil {
			logger.Errorf("unblock user %v", err)
			switch err {
			case repository.ErrUserNotFound:
				ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "user not exist", nil)
				return

			default:
				ginAbortInternalError(c, responsevalue.CodeUnknownError, responsevalue.MsgUnknownError, nil)
				return
			}
		}
		ginRespOK(c, responsevalue.CodeOK, responsevalue.MsgOK, nil)
		return

	default:
		ginAbortNotAcceptable(c, responsevalue.CodeInvalidParameterValue, "invalid block type", nil)
		return
	}
}
