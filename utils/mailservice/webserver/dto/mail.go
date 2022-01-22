package dto

import (
	"github.com/vfluxus/mailservice/repository/entity"
)

type GetMailResp struct {
	RespErr
	Mails []*entity.Mail `json:"mail"`
}

func (r *GetMailResp) SetData(mails ...*entity.Mail) {
	r.Mails = mails
}

type CreateNewMailReq struct {
	FromID     uint32             `json:"from_id"`
	To         []string           `json:"to"`
	Subject    string             `json:"subject"`
	TemplateID uint32             `json:"template_id"`
	Variables  []*entity.Variable `json:"variables"`
}

type CreateNewMailResp struct {
	RespErr
	MailID uint32 `json:"mail_id"`
	HTML   string `json:"html"`
}

func (r *CreateNewMailResp) SetData(mailID uint32, html string) {
	r.MailID = mailID
	r.HTML = html
}

type SendEmailReq struct {
	MailID uint32
}

type SendEmailResp struct {
	RespErr
	HTML string `json:"html"`
}

func (r *SendEmailResp) SetData(html string) {
	r.HTML = html
}

type UpdateMailReq struct {
	MailID    uint32             `json:"mail_id"`
	Mail      *entity.Mail       `json:"mail"`
	Variables []*entity.Variable `json:"variables"`
}

type UpdateMailResp struct {
	RespErr
	HTML string `json:"html"`
}

func (r *UpdateMailResp) SetData(html string) {
	r.HTML = html
}
