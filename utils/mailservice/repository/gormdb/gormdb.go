package gormdb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/vfluxus/mailservice/repository/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// --------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- GORMDB ----------------------------------------------------------

type implGorm struct{}

func (g *implGorm) pkgError(op string, err error) error {
	return fmt.Errorf("pkg: gormdb. Op: %s. Err: %v", op, err)
}

var (
	gDB     = new(gorm.DB)
	gormObj = new(implGorm)
	// model
	mailAccModel  = &entity.MailAccount{}
	templateModel = &entity.Template{}
	mailModel     = &entity.Mail{}
)

const (
	defaultPage = 0
	defaultSize = 10
)

// ------------------------------
// GetGormDB return gorm object to interact with database
func GetGormDB() *implGorm {
	return gormObj
}

// ------------------------------
// SetDBConnection connect gorm to database & set log level(INFO, WANT, ERROR, SILENT)
func (g *implGorm) SetDBConnection(dsn string, logLevel string) (err error) {
	var (
		gormConfig = &gorm.Config{
			Logger: gormlog.Default.LogMode(gormlog.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
	)

	// Change Gorm log level
	switch logLevel {
	case "INFO":
		gormConfig.Logger = gormlog.Default.LogMode(gormlog.Info)
	case "WARN":
		gormConfig.Logger = gormlog.Default.LogMode(gormlog.Warn)
	case "ERROR":
		gormConfig.Logger = gormlog.Default.LogMode(gormlog.Error)
	case "SILENT":
		gormConfig.Logger = gormlog.Default.LogMode(gormlog.Silent)
	default:
		log.Println("START GORM WITH DEFAULT LOG CONFIG: INFO")
	}

	gDB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- AUTO MIGRATE ----------------------------------------------------------

// ------------------------------
// AutoMigrate create tables
func (g *implGorm) AutoMigrate(ctx context.Context) (err error) {
	var (
		models = []interface{}{&entity.MailAccount{}, &entity.Template{}, &entity.Mail{}}
	)

	if err = gDB.WithContext(ctx).AutoMigrate(models...); err != nil {
		return g.pkgError("AutoMigrate", err)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- MAIL ACCOUNT ----------------------------------------------------------

// ------------------------------
// CreateMailAccount ...
func (g *implGorm) CreateMailAccount(ctx context.Context, account *entity.MailAccount) (id uint32, err error) {
	// pre-exec check
	if account == nil {
		return 0, g.pkgError("input check", errors.New("nil accout"))
	}

	if err = gDB.WithContext(ctx).Model(mailAccModel).
		Save(account).Error; err != nil {
		return 0, g.pkgError("CreateMailAccount", err)
	}

	return account.ID, nil
}

// ------------------------------
// SelectAllMailAccounts ...
func (g *implGorm) SelectAllMailAccounts(ctx context.Context, page uint, size uint) (accounts []*entity.MailAccount, err error) {
	// pre-exec check
	page, size = pageSizeCheck(page, size)

	rows, err := gDB.WithContext(ctx).Model(mailAccModel).
		Where("deleted_at IS NULL").
		Order("id ASC").
		Offset(int(page * size)).Limit(int(size)).Rows()
	if err != nil {
		return nil, g.pkgError("SelectAllMailAccounts", err)
	}

	for rows.Next() {
		var acc = new(entity.MailAccount)
		if err = gDB.ScanRows(rows, acc); err != nil {
			return nil, g.pkgError("scan account", err)
		}

		accounts = append(accounts, acc)
	}

	return accounts, nil
}

// ------------------------------
// SelectMailAccountByID ...
func (g *implGorm) SelectMailAccountByID(ctx context.Context, accountID uint32) (account *entity.MailAccount, err error) {
	account = new(entity.MailAccount)

	if err = gDB.WithContext(ctx).Model(mailAccModel).
		Where("id = ? AND deleted_at IS NULL", accountID).
		Take(account).Error; err != nil {
		return nil, g.pkgError("SelectMailAccountByID", err)
	}

	return account, nil
}

// ------------------------------
// UpdateMailAccountByID ...
func (g *implGorm) UpdateMailAccountByID(ctx context.Context, accountID uint32, account *entity.MailAccount) (err error) {
	// pre-exec check
	if account == nil {
		return g.pkgError("input check", errors.New("nil accout"))
	}

	// prevent unexpected udpate
	account.DeletedAt.Valid = false
	account.ID = 0

	if err = gDB.WithContext(ctx).Model(mailAccModel).Where("id = ?", accountID).Updates(account).Error; err != nil {
		return g.pkgError("UpdateMailAccountByID", err)
	}
	return
}

// ------------------------------
// DeleteMailAccountByID ...
func (g *implGorm) DeleteMailAccountByID(ctx context.Context, accountID uint32) (err error) {
	if err = gDB.WithContext(ctx).Model(mailAccModel).Where("id = ?", accountID).Update("deleted_at", time.Now()).Error; err != nil {
		return g.pkgError("DeleteMailAccountByID", err)
	}

	return nil
}

// ----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- TEMPLATE ----------------------------------------------------------

// ------------------------------
// CreateTemplate ...
func (g *implGorm) CreateTemplate(ctx context.Context, template *entity.Template) (id uint32, err error) {
	// pre-exec check
	if template == nil {
		return 0, g.pkgError("input check", errors.New("nil template"))
	}

	if err = gDB.WithContext(ctx).Model(templateModel).Save(template).Error; err != nil {
		return 0, g.pkgError("CreateTemplate", err)
	}

	return template.ID, nil
}

// ------------------------------
// SelectAllTemplates ...
func (g *implGorm) SelectAllTemplates(ctx context.Context, page uint, size uint) (templates []*entity.Template, err error) {
	// pre-exec check
	page, size = pageSizeCheck(page, size)

	rows, err := gDB.WithContext(ctx).Model(templateModel).
		Where("deleted_at IS NULL").
		Order("id ASC").
		Offset(int(page * size)).Limit(int(size)).Rows()
	if err != nil {
		return nil, g.pkgError("SelectAllTemplates", err)
	}

	for rows.Next() {
		var template = new(entity.Template)

		if err = gDB.WithContext(ctx).ScanRows(rows, template); err != nil {
			return nil, g.pkgError("scan template", err)
		}

		templates = append(templates, template)
	}

	return templates, nil
}

// ------------------------------
// SelectTemplateByID ...
func (g *implGorm) SelectTemplateByID(ctx context.Context, templateID uint32) (template *entity.Template, err error) {
	template = new(entity.Template)

	if err = gDB.WithContext(ctx).Model(templateModel).
		Where("id = ? AND deleted_at IS NULL", templateID).
		Take(template).Error; err != nil {
		return nil, g.pkgError("SelectTemplateByID", err)
	}

	return template, nil
}

func (g *implGorm) SelectTemplateByMailID(ctx context.Context, mailID uint32) (template *entity.Template, err error) {
	template = new(entity.Template)

	err = gDB.WithContext(ctx).Select("template.*").
		Model(templateModel).
		Joins("JOIN mail ON mail.template_id = template.id").
		Where("mail.ID = ?", mailID).
		Take(template).Error

	if err != nil {
		return nil, err
	}

	return template, nil
}

// ------------------------------
// UpdateTemplateByID ...
func (g *implGorm) UpdateTemplateByID(ctx context.Context, templateID uint32, template *entity.Template) (err error) {
	// pre-exec check
	if template == nil {
		return g.pkgError("input check", errors.New("nil template"))
	}
	// prevent unexpected update
	template.ID = 0
	template.DeletedAt.Valid = false

	if err = gDB.WithContext(ctx).Model(templateModel).Where("id = ?", templateID).Updates(template).Error; err != nil {
		return g.pkgError("UpdateTemplateByID", err)
	}

	return nil
}

// ------------------------------
// DeleteTemplateByID ...
func (g *implGorm) DeleteTemplateByID(ctx context.Context, templateID uint32) (err error) {
	if err = gDB.WithContext(ctx).Model(templateModel).Where("id = ?", templateID).Update("deleted_at", time.Now()).Error; err != nil {
		return g.pkgError("DeleteTemplateByID", err)
	}

	return nil
}

// ------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- MAIL ----------------------------------------------------------

// ------------------------------
// CreateMail ...
func (g *implGorm) CreateMail(ctx context.Context, mail *entity.Mail) (id uint32, err error) {
	// pre-exec check
	if mail == nil {
		return 0, g.pkgError("input check", errors.New("nil mail"))
	}

	if err = gDB.WithContext(ctx).Model(mailModel).Save(mail).Error; err != nil {
		return 0, g.pkgError("CreateMail", err)
	}

	return mail.ID, nil
}

// ------------------------------
// SelectAllMail ...
func (g *implGorm) SelectAllMail(ctx context.Context, page uint, size uint) (mails []*entity.Mail, err error) {
	page, size = pageSizeCheck(page, size)

	err = gDB.WithContext(ctx).Model(mailModel).
		Order("id ASC").
		Offset(int(page * size)).Limit(int(size)).
		Find(&mails).Error

	if err != nil {
		return nil, err
	}

	return mails, nil
}

// ------------------------------
// SelectMailByID ...
func (g *implGorm) SelectMailByID(ctx context.Context, mailID uint32) (mail *entity.Mail, err error) {
	mail = new(entity.Mail)

	if err = gDB.WithContext(ctx).Model(mailModel).
		Where("id = ? AND deleted_at IS NULL", mailID).
		Take(mail).Error; err != nil {
		return nil, g.pkgError("SelectMailByID", err)
	}
	return
}

// ------------------------------
// SelectMailsSentByAccountID ...
func (g *implGorm) SelectMailsSentByAccountID(ctx context.Context, accountID uint32, page uint, size uint) (mails []*entity.Mail, err error) {
	// pre-exec check
	page, size = pageSizeCheck(page, size)

	rows, err := gDB.WithContext(ctx).Model(mailModel).
		Where("account_id = ? AND deleted_at IS NULL", accountID).
		Order("id ASC").
		Offset(int(page * size)).Limit(int(size)).Rows()

	if err != nil {
		return nil, g.pkgError("SelectMailsSentByAccountID", err)
	}

	for rows.Next() {
		var mail = new(entity.Mail)
		if err = gDB.WithContext(ctx).ScanRows(rows, mail); err != nil {
			return nil, g.pkgError("scan mail", err)
		}

		mails = append(mails, mail)
	}

	return mails, nil
}

// ------------------------------
// SelectMailsSentByTemplateID ...
func (g *implGorm) SelectMailsSentByTemplateID(ctx context.Context, templateID uint32, page uint, size uint) (mails []*entity.Mail, err error) {
	// pre-exec check
	page, size = pageSizeCheck(page, size)

	rows, err := gDB.WithContext(ctx).Model(mailModel).
		Where("template_id = ? AND deleted_at IS NULL", templateID).
		Order("id ASC").
		Offset(int(page * size)).Limit(int(size)).Rows()

	if err != nil {
		return nil, g.pkgError("SelectMailsSentByAccountID", err)
	}

	for rows.Next() {
		var mail = new(entity.Mail)
		if err = gDB.WithContext(ctx).ScanRows(rows, mail); err != nil {
			return nil, g.pkgError("scan mail", err)
		}

		mails = append(mails, mail)
	}

	return mails, nil
}

// ------------------------------
// UpdateMailByID ...
func (g *implGorm) UpdateMailByID(ctx context.Context, mailID uint32, mail *entity.Mail) (err error) {
	// pre-exec check
	if mail == nil {
		return g.pkgError("input check", errors.New("nil mail"))
	}
	// prevent unexpected update
	mail.ID = 0

	if err = gDB.WithContext(ctx).Model(mailModel).
		Where("id = ? AND deleted_at IS NULL", mailID).
		Updates(mail).Error; err != nil {
		return g.pkgError("UpdateMailByID", err)
	}

	return
}

// ------------------------------
// DeleteMailByID ...
func (g *implGorm) DeleteMailByID(ctx context.Context, mailID uint32) (err error) {
	if err = gDB.WithContext(ctx).Model(mailModel).
		Where("id = ? AND deleted_at IS NULL", mailID).
		Update("deleted_at", time.Now()).Error; err != nil {
		return g.pkgError("DeleteMailByID", err)
	}

	return
}

// -------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- UTILS ----------------------------------------------------------

func pageSizeCheck(page uint, size uint) (newPage uint, newSize uint) {
	if page < 1 {
		newPage = defaultPage
	} else {
		newPage = page - 1
	}

	if size == 0 {
		newSize = defaultSize
	}

	return newPage, newSize
}
