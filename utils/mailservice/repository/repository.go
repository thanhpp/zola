package repository

import (
	"context"

	"github.com/vfluxus/mailservice/repository/entity"
	"github.com/vfluxus/mailservice/repository/gormdb"
)

type iDAO interface {
	SetDBConnection(dsn string, logLevel string) (err error)
	AutoMigrate(ctx context.Context) (err error)

	CreateMailAccount(ctx context.Context, account *entity.MailAccount) (id uint32, err error)
	SelectAllMailAccounts(ctx context.Context, page uint, size uint) (accounts []*entity.MailAccount, err error)
	SelectMailAccountByID(ctx context.Context, accountID uint32) (account *entity.MailAccount, err error)
	UpdateMailAccountByID(ctx context.Context, accountID uint32, account *entity.MailAccount) (err error)
	DeleteMailAccountByID(ctx context.Context, accountID uint32) (err error)

	CreateTemplate(ctx context.Context, template *entity.Template) (id uint32, err error)
	SelectAllTemplates(ctx context.Context, page uint, size uint) (templates []*entity.Template, err error)
	SelectTemplateByID(ctx context.Context, templateID uint32) (template *entity.Template, err error)
	SelectTemplateByMailID(ctx context.Context, mailID uint32) (template *entity.Template, err error)
	UpdateTemplateByID(ctx context.Context, templateID uint32, template *entity.Template) (err error)
	DeleteTemplateByID(ctx context.Context, templateID uint32) (err error)

	CreateMail(ctx context.Context, mail *entity.Mail) (id uint32, err error)
	SelectAllMail(ctx context.Context, page uint, size uint) (mails []*entity.Mail, err error)
	SelectMailByID(ctx context.Context, mailID uint32) (mail *entity.Mail, err error)
	SelectMailsSentByAccountID(ctx context.Context, accountID uint32, page uint, size uint) (mails []*entity.Mail, err error)
	SelectMailsSentByTemplateID(ctx context.Context, templateID uint32, page uint, size uint) (mails []*entity.Mail, err error)
	UpdateMailByID(ctx context.Context, mailID uint32, mail *entity.Mail) (err error)
	DeleteMailByID(ctx context.Context, mailID uint32) (err error)
}

func GetDAO() iDAO {
	return gormdb.GetGormDB()
}
