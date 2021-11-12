package entity

import "errors"

var (
	ReportNil = Report{}
)

var (
	ErrReportLockedPost = errors.New("report locked post")
)

type Report struct {
	ID        string
	PostID    string
	Detail    string
	CreatedBy string
	Subject   ReportSubject
}

func (r Report) GetID() string {
	return r.ID
}

func (r Report) GetPostID() string {
	return r.PostID
}

func (r Report) GetDetail() string {
	return r.Detail
}

func (r Report) GetUserID() string {
	return r.CreatedBy
}

func (r Report) GetSubject() string {
	return r.Subject.Subject
}

func (r Report) GetSubjectID() int {
	return r.Subject.ID
}
