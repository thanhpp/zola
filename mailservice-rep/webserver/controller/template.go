package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vfluxus/dvergr/logger"
	"github.com/vfluxus/mailservice/repository"
	"github.com/vfluxus/mailservice/service"
	"github.com/vfluxus/mailservice/webserver/dto"
)

type TemplateController struct{}

// ------------------------------
// GetTemplate ...
// @Summary Get template
// @Description Get template, if id is not specified, then get all template with page and size
// @Produce json
// @Param 	id 		query	int		false 	"template id"
// @Param 	page	query	int		false 	"page"
// @Param	size	query	int		false	"size"
// @Success 200 	{object}	dto.GetTemplateResp	"Get OK"
// @Tags template
// @Router /template [GET]
func (t *TemplateController) GetTemplate(c *gin.Context) {
	id, err := getIDFromQuery(c)
	if err != nil && !errors.Is(err, errEmptyID) {
		logger.Get().Errorf("Get param id error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	// if id is found, then get by id
	if !errors.Is(err, errEmptyID) {
		template, err := repository.GetDAO().SelectTemplateByID(c, id)
		if err != nil {
			logger.Get().Errorf("GetTemplate DB err: %v", err)
			ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
			return
		}

		resp := new(dto.GetTemplateResp)
		resp.SetCode(http.StatusOK)
		resp.SetData(template)
		c.JSON(http.StatusOK, resp)
		return
	}

	page, size, err := getPageSizeFromQuery(c)
	if err != nil {
		logger.Get().Errorf("GetTemplate get page & size error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	templates, err := repository.GetDAO().SelectAllTemplates(c, page, size)
	if err != nil {
		logger.Get().Errorf("GetTemplate DB error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.GetTemplateResp)
	resp.SetCode(http.StatusOK)
	resp.SetData(templates...)
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// CreateNewTemplate
// @Summary Create new template
// @Description Create new template
// @Produce json
// @Param createReq		body	dto.CreateNewTemplateReq	true	"Create info"
// @Success 200		{object}	dto.CreateNewTemplateResp	true	"Create OK"
// @Tags template
// @Router /template [POST]
func (t *TemplateController) CreateNewTemplate(c *gin.Context) {
	req := new(dto.CreateNewTemplateReq)
	if err := c.ShouldBindJSON(req); err != nil {
		logger.Get().Errorf("CreateNewTemplate bind json error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	if err := mergeTemplate(req.Template, req.Variables); err != nil {
		logger.Get().Errorf("CreateNewTemplate merge template error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	// validate template
	if _, err := service.GetTemplate().Validate(req.Template); err != nil {
		logger.Get().Errorf("CreateNewTemplate validate template error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := repository.GetDAO().CreateTemplate(c, req.Template)
	if err != nil {
		logger.Get().Errorf("CreateNewTemplate DB error: %v", err)
		ginRespErrAbort(c, http.StatusOK, err.Error())
		return
	}

	resp := new(dto.CreateNewTemplateResp)
	resp.SetCode(http.StatusOK)
	resp.SetData(id)
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// UpdateTemplate ...
// @Summary Update template
// @Description Update template, need to specified templateID, will not validate template
// @Produce json
// @Param 	updateReq	body	dto.UpdateTemplateReq	true	"update req"
// @Success 200 	{object}	dto.RespErr		"Update OK"
// @Tags template
// @Router /template [PUT]
func (t *TemplateController) UpdateTemplate(c *gin.Context) {
	req := new(dto.UpdateTemplateReq)
	if err := c.ShouldBindJSON(req); err != nil {
		logger.Get().Errorf("UpdateTemplate bind json error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	if err := mergeTemplate(req.Template, req.Variables); err != nil {
		logger.Get().Errorf("UpdateTemplate merge template error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := repository.GetDAO().UpdateTemplateByID(c, req.TemplateID, req.Template); err != nil {
		logger.Get().Errorf("UpdateTemplate DB error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.RespErr)
	resp.SetCode(http.StatusOK)
	c.JSON(http.StatusOK, resp)
}

// ------------------------------
// DeleteTemplate ...
// @Summary DeleteTemplate
// @Description DeleteTemplate by id
// @Produce json
// @Param 	id	query	int		true	"templateID"
// @Success 200	{object}	dto.RespErr	"DeleteOK"
// @Tags template
// @Router /template [DELETE]
func (t *TemplateController) DeleteTemplate(c *gin.Context) {
	id, err := getIDFromQuery(c)
	if err != nil {
		logger.Get().Errorf("DeleteTemplate Get param id error: %v", err)
		ginRespErrAbort(c, http.StatusNotAcceptable, err.Error())
		return
	}

	if err := repository.GetDAO().DeleteTemplateByID(c, id); err != nil {
		logger.Get().Errorf("DeleteTemplate DB error: %v", err)
		ginRespErrAbort(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := new(dto.RespErr)
	resp.SetCode(http.StatusOK)
	c.JSON(http.StatusOK, resp)
}
