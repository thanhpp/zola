package app

import "errors"

var (
	ErrBlocked      = errors.New("blocked")
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidMsgID = errors.New("invalid msgID")
	ErrNotSender    = errors.New("not sender")
)
