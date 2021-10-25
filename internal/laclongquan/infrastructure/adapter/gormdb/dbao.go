package gormdb

import "github.com/thanhpp/zola/config/shared"

type DBAO struct {
	User userGorm
}

func NewDBAO(cfg *shared.DatabaseConfig) (*DBAO, error) {
	if err := initConnection(cfg.DSN(), cfg.LogLevel, cfg.LogColor); err != nil {
		return nil, err
	}

	return &DBAO{
		User: userGorm{
			db:    gDB,
			model: &UserDB{},
		},
	}, nil
}
