package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vfluxus/dvergr/logger"
	"github.com/vfluxus/mailservice/repository"
	"github.com/vfluxus/mailservice/webserver/dto"
)

type AccountController struct{}

// ------------------------------
// GetAccount
// @Summary Get account
// @Description Get account if id is not specified, then get all account with page and size
// @Produce json
// @Param 	id 		query	int		false 	"account id"
// @Param 	page	query	int		false 	"page"
// @Param	size	query	int		false	"size"
// @Success 200		{object}	dto.GetMailAccountResp 	"Get success"
// @Tags account
// @Router /account [GET]
func (a *AccountController) GetAccount(c *gin.Context) {
	id, err := getIDFromQuery(c)
	if err != nil && !errors.Is(err, errEmptyID) {
		logger.Get().Errorf("Get param id error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	// if no id found, then get all
	if !errors.Is(err, errEmptyID) {
		account, err := repository.GetDAO().SelectMailAccountByID(c, id)
		if err != nil {
			logger.Get().Errorf("Select account by id error: %v", err)
			ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
			return
		}

		logger.Get().Debugf("Get mail id :%d OK", id)
		resp := new(dto.GetMailAccountResp)
		resp.SetCode(http.StatusOK)
		resp.SetData(account)
		c.JSON(http.StatusOK, resp)
		return
	}

	page, size, err := getPageSizeFromQuery(c)
	if err != nil {
		logger.Get().Errorf("get page & size error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	accounts, err := repository.GetDAO().SelectAllMailAccounts(c, page, size)
	if err != nil {
		logger.Get().Errorf("get all mail error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Get().Debugf("Get all mail page %d, size %d OK", page, size)
	resp := new(dto.GetMailAccountResp)
	resp.SetCode(http.StatusOK)
	resp.SetData(accounts...)
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// CreateNewAccount ...
// @Summary Create new account
// @Description Create new account
// @Produce json
// @Param createReq 	body 	dto.CreateMailAccountReq	true "account info"
// @Success 200 	{object} 	dto.CreateMailAccountResp	"Create OK"
// @Tags account
// @Router /account [POST]
func (a *AccountController) CreateNewAccount(c *gin.Context) {
	req := new(dto.CreateMailAccountReq)
	if err := c.ShouldBindJSON(req); err != nil {
		logger.Get().Errorf("CreateNewAccount - bind json error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	hashedPass, err := encryptPassword(req.Account.Password)
	if err != nil {
		logger.Get().Errorf("CreateNewAccount - Encrypt pass error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}
	fmt.Println(hashedPass)
	req.Account.Password = hashedPass

	accountID, err := repository.GetDAO().CreateMailAccount(c, req.Account)
	if err != nil {
		logger.Get().Errorf("CreateMailAccount DB err: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Get().Debug("CreateNewAccount OK")
	resp := new(dto.CreateMailAccountResp)
	resp.SetCode(http.StatusOK)
	resp.SetData(accountID)
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// UpdateAccountInfo ...
// @Summary Update account info
// @Description Update account info (Need to specified account ID)
// @Produce json
// @Param 	updateReq	body	dto.UpdateMailAccountReq	true	"Update info"
// @Success 200 	{object}	dto.RespErr		"Update OK"
// @Tags account
// @Router /account [PUT]
func (a *AccountController) UpdateAccountInfo(c *gin.Context) {
	req := new(dto.UpdateMailAccountReq)
	if err := c.ShouldBindJSON(req); err != nil {
		logger.Get().Errorf("UpdateAccountInfo bind jsone err: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	if req.Account != nil && len(req.Account.Password) > 0 {
		hashedPass, err := encryptPassword(req.Account.Password)
		if err != nil {
			logger.Get().Errorf("UpdateAccountInfo - Encrypt pass error: %v", err)
			ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
			return
		}

		req.Account.Password = hashedPass
	}

	if err := repository.GetDAO().UpdateMailAccountByID(c, req.AccountID, req.Account); err != nil {
		logger.Get().Errorf("UpdateAccountInfo DB err: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Get().Debug("UpdateAccountInfo OK")
	resp := new(dto.RespErr)
	resp.SetCode(http.StatusOK)
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// DeleteAccount ...
// @Summary Delete account by id
// @Description Delete account by id
// @Produce json
// @Param 	id	query	int		true	"account id"
// @Success 200	{object} 	dto.RespErr	"delete OK"
// @Tags account
// @Router /account [DELETE]
func (a *AccountController) DeleteAccount(c *gin.Context) {
	id, err := getIDFromQuery(c)
	if err != nil {
		logger.Get().Errorf("DeleteAccount Get param id error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	if err = repository.GetDAO().DeleteMailAccountByID(c, id); err != nil {
		logger.Get().Errorf("DeleteAccount DB error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.RespErr)
	resp.SetCode(http.StatusOK)
	c.JSON(http.StatusOK, resp)
}
