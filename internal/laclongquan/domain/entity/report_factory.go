package entity

import "github.com/google/uuid"

type ReportFactory interface {
	NewReport(postID, userID string, subjectID int, detail string) (Report, error)
}

func NewReportFactory() ReportFactory {
	return reportFactoryImpl{}
}

type reportFactoryImpl struct{}

func (f reportFactoryImpl) NewReport(postID, creator string, subjectID int, detail string) (Report, error) {
	subject, err := reportSubjectFromID(subjectID)
	if err != nil {
		return ReportNil, err
	}

	return Report{
		ID:        uuid.NewString(),
		PostID:    postID,
		CreatedBy: creator,
		Subject:   subject,
		Detail:    detail,
	}, nil
}
