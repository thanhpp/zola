package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl UserController) Signout(c *gin.Context) {
	claims, err := getClaimsFromCtx(c)
	if err != nil {
		logger.Errorf("get claims %v", err)
		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	if err := ctrl.authsrv.DeleteUserTokens(c, claims.User.ID); err != nil {
		logger.Errorf("delete user tokens %v", err)
		ginAbortInternalError(c, responsevalue.ValueUnknownError, responsevalue.MsgUnknownError)
		return
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}
