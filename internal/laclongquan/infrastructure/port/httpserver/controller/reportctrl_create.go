package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/dto"
	"github.com/thanhpp/zola/pkg/logger"
	"github.com/thanhpp/zola/pkg/responsevalue"
)

func (ctrl ReportController) Create(c *gin.Context) {
	var req = new(dto.CreateReportReq)

	if err := c.ShouldBind(req); err != nil {
		logger.Errorf("bind req err: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, responsevalue.MsgInvalidRequest)
		return
	}

	// get creator id
	userID, err := getUserUUIDFromClaims(c)
	if err != nil {
		logger.Errorf("get user id from ctx err: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidateUser, responsevalue.MsgInvalidRequest)
		return
	}

	// convert subject to id
	subjectInt, err := strconv.Atoi(req.Subject)
	if err != nil {
		logger.Errorf("convert subject to id err: %v", err)
		ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, responsevalue.MsgInvalidRequest)
		return
	}

	_, err = ctrl.app.CreateReport(c, req.ID, userID.String(), subjectInt, req.Details)
	if err != nil {
		logger.Errorf("create report err: %v", err)
		switch err {
		case entity.ErrInvalidReportSubjectID:
			ginAbortNotAcceptable(c, responsevalue.ValueInvalidParameterType, "invalid subject id")
			return

		case repository.ErrPostNotFound:
			ginAbortNotAcceptable(c, responsevalue.ValuePostNotExist, "post not exist")
			return

		case entity.ErrLockedPost:
			ginAbortNotAcceptable(c, responsevalue.ValueActionHasBeenDone, "report locked post")
			return
		}

		ginAbortInternalError(c, responsevalue.ValueUnknownError, err.Error())
		return
	}

	ginRespOK(c, responsevalue.ValueOK, responsevalue.MsgOK)
}
