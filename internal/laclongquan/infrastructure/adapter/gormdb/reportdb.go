package gormdb

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"gorm.io/gorm"
)

type ReportDB struct {
	ReportUUID  string `gorm:"Column:report_uuid; Type:text; primaryKey"`
	PostUUID    string `gorm:"Column:post_uuid; Type:text"`
	CreatorUUID string `gorm:"Column:creator_uuid; Type:text"`
	SubjectID   int    `gorm:"Column:subject_id; Type:bigint"`
	Detail      string `gorm:"Column:detail; Type:text"`
}

type reportGorm struct {
	db    *gorm.DB
	model *ReportDB
}

func (r reportGorm) marshalReport(report *entity.Report) *ReportDB {
	return &ReportDB{
		ReportUUID:  report.GetID(),
		PostUUID:    report.GetPostID(),
		CreatorUUID: report.GetUserID(),
		SubjectID:   report.GetSubjectID(),
		Detail:      report.GetDetail(),
	}
}

func (r reportGorm) Create(ctx context.Context, report *entity.Report) error {
	reportDB := r.marshalReport(report)

	return r.db.WithContext(ctx).Model(r.model).Create(reportDB).Error
}
