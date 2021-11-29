package gormdb

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	gDB = new(gorm.DB)
)

func initConnection(dsn string, logLevel string, color bool) error {
	gormCfg := &gorm.Config{
		Logger: gormlog.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormlog.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      gormlog.Info,
				Colorful:      color,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	switch strings.ToLower(logLevel) {
	case "info":
		break
	case "warn":
		gormCfg.Logger = gormlog.Default.LogMode(gormlog.Warn)
	case "error":
		gormCfg.Logger = gormlog.Default.LogMode(gormlog.Error)
	case "silent":
		gormCfg.Logger = gormlog.Default.LogMode(gormlog.Silent)
	}

	newGormDB, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		return err
	}

	gDB = newGormDB

	return autoMigrate()
}

func autoMigrate() error {
	return gDB.AutoMigrate(
		// user
		&UserDB{},
		&AuthDB{},
		&RelationDB{},
		// post
		&PostDB{},
		&MediaDB{},
		&ReportDB{},
		&LikeDB{},
		&CommentDB{},
	)
}

func isDuplicate(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return true
		}
	}
	return false
}
