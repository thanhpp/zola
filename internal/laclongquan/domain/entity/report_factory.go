package entity

import "github.com/google/uuid"

type ReportFactory interface {
	NewReport(post *Post, userID string, subjectID int, detail string) (Report, error)
}

func NewReportFactory() ReportFactory {
	return reportFactoryImpl{}
}

type reportFactoryImpl struct{}

func (f reportFactoryImpl) NewReport(post *Post, creator string, subjectID int, detail string) (Report, error) {
	if post.IsLocked() {
		return ReportNil, ErrReportLockedPost
	}

	subject, err := reportSubjectFromID(subjectID)
	if err != nil {
		return ReportNil, err
	}

	return Report{
		ID:        uuid.NewString(),
		PostID:    post.ID(),
		CreatedBy: creator,
		Subject:   subject,
		Detail:    detail,
	}, nil
}
