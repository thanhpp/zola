package dto

import (
	"github.com/vfluxus/mailservice/repository/entity"
)

type GetMailAccountResp struct {
	RespErr
	Accounts []*entity.MailAccount `json:"accounts"`
}

func (r *GetMailAccountResp) SetData(account ...*entity.MailAccount) {
	r.Accounts = account
}

type CreateMailAccountReq struct {
	Account *entity.MailAccount `json:"account"`
}

type CreateMailAccountResp struct {
	RespErr
	AccountID uint32 `json:"account_id"`
}

func (r *CreateMailAccountResp) SetData(accountID uint32) {
	r.AccountID = accountID
}

type UpdateMailAccountReq struct {
	AccountID uint32              `json:"account_id"`
	Account   *entity.MailAccount `json:"account"`
}
