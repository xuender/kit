package times

import "errors"

var (
	ErrLast       = errors.New("last id after now")
	ErrParseError = errors.New("parse error")
)
