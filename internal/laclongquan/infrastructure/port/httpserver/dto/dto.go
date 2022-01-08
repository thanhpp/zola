package dto

import "errors"

var (
	ErrInvalidBoolString = errors.New("invalid bool string")
)

func boolTranslate(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func BoolTranslateStr(str string) (bool, error) {
	switch str {
	case "1":
		return true, nil
	case "0":
		return false, nil
	default:
		return false, ErrInvalidBoolString
	}
}
