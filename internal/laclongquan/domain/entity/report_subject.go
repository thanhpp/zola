package entity

import "errors"

type ReportSubject struct {
	ID      int
	Subject string
}

var (
	ErrInvalidReportSubjectID = errors.New("invalid report subject id")
)

var (
	ReportSubjectNil              = ReportSubject{}
	ReportSubjectOther            = ReportSubject{ID: 0, Subject: "Other"}
	ReportSubjectSensitiveContent = ReportSubject{ID: 1, Subject: "Sensitive Content"}
	ReportSubjectDisturbing       = ReportSubject{ID: 2, Subject: "Disturbing"}
	ReportSubjectScam             = ReportSubject{ID: 3, Subject: "Scam"}
)

func (r ReportSubject) GetSubject() string {
	return r.Subject
}

func reportSubjectFromID(id int) (ReportSubject, error) {
	switch id {
	case 0:
		return ReportSubjectOther, nil

	case 1:
		return ReportSubjectSensitiveContent, nil

	case 2:
		return ReportSubjectDisturbing, nil

	case 3:
		return ReportSubjectScam, nil

	default:
		return ReportSubjectNil, ErrInvalidReportSubjectID
	}
}
