package dto

import (
	"github.com/vfluxus/mailservice/repository/entity"
)

type GetTemplateResp struct {
	RespErr
	Templates []*entity.Template `json:"template"`
}

func (r *GetTemplateResp) SetData(templates ...*entity.Template) {
	r.Templates = templates
}

type CreateNewTemplateReq struct {
	Template  *entity.Template   `json:"template"`
	Variables []*entity.Variable `json:"variables"`
}

type CreateNewTemplateResp struct {
	RespErr
	TemplateID uint32 `json:"template_id"`
}

func (r *CreateNewTemplateResp) SetData(templateID uint32) {
	r.TemplateID = templateID
}

type UpdateTemplateReq struct {
	TemplateID uint32
	Template   *entity.Template   `json:"template"`
	Variables  []*entity.Variable `json:"variables"`
}
