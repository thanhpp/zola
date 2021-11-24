package gormdb

import "github.com/thanhpp/zola/config/shared"

type DBAO struct {
	User     userGorm
	Auth     authGorm
	Post     postGorm
	Report   reportGorm
	Like     likeGorm
	Relation relationGorm
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
		Auth: authGorm{
			db:    gDB,
			model: &AuthDB{},
		},
		Post: postGorm{
			db:         gDB,
			mediaModel: &MediaDB{},
			postModel:  &PostDB{},
		},
		Report: reportGorm{
			db:    gDB,
			model: &ReportDB{},
		},
		Like: likeGorm{
			db:    gDB,
			model: &LikeDB{},
		},
		Relation: relationGorm{
			db:    gDB,
			model: &RelationDB{},
		},
	}, nil
}
