package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vfluxus/dvergr/logger"
	"github.com/vfluxus/mailservice/repository"
	"github.com/vfluxus/mailservice/repository/entity"
	"github.com/vfluxus/mailservice/service"
	"github.com/vfluxus/mailservice/webserver/dto"
)

type MailController struct{}

// ------------------------------
// GetMail ...
// @Summary Get mail
// @Description Get mail, if id is not specified, get all by page and size
// @Produce json
// @Param 	id 		query	int		false 	"mail id"
// @Param 	page	query	int		false 	"page"
// @Param	size	query	int		false	"size"
// @Success	200		{object}	dto.GetMailResp		"Get OK"
// @Tags mail
// @Router /mail [GET]
func (m *MailController) GetMail(c *gin.Context) {
	id, err := getIDFromQuery(c)
	if err != nil && !errors.Is(err, errEmptyID) {
		logger.Get().Errorf("GetMail get id error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	// found id
	if !errors.Is(err, errEmptyID) {
		mail, err := repository.GetDAO().SelectMailByID(c, id)
		if err != nil {
			logger.Get().Errorf("GetMail DB error: %v", err)
			ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
			return
		}

		resp := new(dto.GetMailResp)
		resp.SetCode(http.StatusOK)
		resp.SetData(mail)
		c.JSON(http.StatusOK, resp)
		return
	}

	page, size, err := getPageSizeFromQuery(c)
	if err != nil {
		logger.Get().Errorf("GetMail get page size error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	mails, err := repository.GetDAO().SelectAllMail(c, page, size)
	if err != nil {
		logger.Get().Errorf("GetMail DB error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.GetMailResp)
	resp.SetCode(http.StatusOK)
	resp.SetData(mails...)
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// CreateNewMail ...
// @Summary Create new mail
// @Description Create new mail, will generate html to preview
// @Produce json
// @Param 	createReq	body	dto.CreateNewMailReq	true	"create req"
// @Success 200		{object}	dto.CreateNewMailResp	"Create OK"
// @Tags mail
// @Router /mail [POST]
func (m *MailController) CreateNewMail(c *gin.Context) {
	req := new(dto.CreateNewMailReq)
	if err := c.ShouldBindJSON(req); err != nil {
		logger.Get().Errorf("CreateNewMail bind json error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	// get template
	template, err := repository.GetDAO().SelectTemplateByID(c, req.TemplateID)
	if err != nil {
		logger.Get().Errorf("CreateNewMail DB get template error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	// parse html
	html, err := service.GetTemplate().Parse(template, req.Variables)
	if err != nil {
		logger.Get().Errorf("CreateNewMail parse html error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	// pass variable to mail
	data, err := json.Marshal(req.Variables)
	if err != nil {
		logger.Get().Errorf("CreateNewMail marshal vars error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	// save to DB
	mail := &entity.Mail{
		AccountID:  req.FromID,
		SendTo:     req.To,
		Subject:    req.Subject,
		TemplateID: req.TemplateID,
		Varibles:   data,
		Status:     "UNSENT",
	}
	id, err := repository.GetDAO().CreateMail(c, mail)
	if err != nil {
		logger.Get().Errorf("CreateNewMail DB error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.CreateNewMailResp)
	resp.SetCode(http.StatusOK)
	resp.SetData(id, string(html))
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// SendMail ...
// @Summary Send email
// @Description Send email
// @Produce json
// @Param 	SendReq		body	dto.SendEmailReq	true 	"mail info"
// @Success 200		{object}	dto.SendEmailResp	"Send OK"
// @Tags 	mail
// @Router /mail/send [POST]
func (m *MailController) SendMail(c *gin.Context) {
	req := new(dto.SendEmailReq)
	if err := c.ShouldBindJSON(req); err != nil {
		logger.Get().Errorf("SendMail bind json error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	mail, err := repository.GetDAO().SelectMailByID(c, req.MailID)
	if err != nil {
		logger.Get().Errorf("SendMail DB get mail rror: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	// get & parse template
	template, err := repository.GetDAO().SelectTemplateByID(c, mail.TemplateID)
	if err != nil {
		logger.Get().Errorf("SendMail DB get template err: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}
	var vars []*entity.Variable
	if err := json.Unmarshal(mail.Varibles, &vars); err != nil {
		logger.Get().Errorf("SendMail unmarshal vars error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}
	html, err := service.GetTemplate().Parse(template, vars)
	if err != nil {
		logger.Get().Errorf("SendMail parse template error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	// get account info
	account, err := repository.GetDAO().SelectMailAccountByID(c, mail.AccountID)
	if err != nil {
		logger.Get().Errorf("SendMail DB get account error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	// decrypt password
	password, err := decryptPassword(account.Password)
	if err != nil {
		logger.Get().Errorf("SendMail decrypt password error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}
	account.Password = password

	// send email
	if err := service.GetSMTP().SendEmail(account, mail, html); err != nil {
		logger.Get().Errorf("SendEmail error: %v", err)
		mail.Status = err.Error()
		if err := repository.GetDAO().UpdateMailByID(c, mail.ID, mail); err != nil {
			logger.Get().Errorf("SendEmail Update mail status error: %v", err)
			ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
			return
		}
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	mail.Status = "SENT"
	mail.SentAt = time.Now()

	if err := repository.GetDAO().UpdateMailByID(c, mail.ID, mail); err != nil {
		logger.Get().Errorf("SendEmail Update mail status error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.SendEmailResp)
	resp.SetCode(http.StatusOK)
	resp.SetData(string(html))
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// UpdateMail ...
// @Summary update email
// @Description update email, need to specify id, update variables will update entire mail variables
// @Produce json
// @Param 	updateReq	body	dto.UpdateMailReq	true	"Update info"
// @Success	200		{object}	dto.UpdateMailResp	"Update OK"
// @Tags	mail
// @Router /mail [PUT]
func (m *MailController) UpdateMail(c *gin.Context) {
	req := new(dto.UpdateMailReq)
	if err := c.ShouldBindJSON(req); err != nil {
		logger.Get().Errorf("UpdateMail bind json error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	template, err := repository.GetDAO().SelectTemplateByMailID(c, req.MailID)
	if err != nil {
		logger.Get().Errorf("UpdateMail DB get template error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	html, err := service.GetTemplate().Parse(template, req.Variables)
	if err != nil {
		logger.Get().Errorf("UpdateMail parse html error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := json.Marshal(req.Variables)
	if err != nil {
		logger.Get().Errorf("UpdateMail marshal variables error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}
	if req.Mail == nil && req.Variables != nil {
		req.Mail = new(entity.Mail) // prevent panic
		req.Mail.Varibles = data
	}

	if err := repository.GetDAO().UpdateMailByID(c, req.MailID, req.Mail); err != nil {
		logger.Get().Errorf("UpdateMail DB error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.UpdateMailResp)
	resp.SetCode(http.StatusOK)
	resp.SetData(string(html))
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// DeleteMail ...
// @Summary Delete mail
// @Description Delete mail by id
// @Produce 	json
// @Param 	id	query	int		true	"mail ID"
// @Success 200	{object}	dto.RespErr	"delete OK"
// @Tags mail
// @Router /delete [DELETE]
func (m *MailController) DeleteMail(c *gin.Context) {
	id, err := getIDFromQuery(c)
	if err != nil {
		logger.Get().Errorf("DeleteMail get id param error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	if err := repository.GetDAO().DeleteMailByID(c, id); err != nil {
		logger.Get().Errorf("DeleteMail DB error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.RespErr)
	resp.SetCode(http.StatusOK)
	c.JSON(http.StatusOK, resp)
}
