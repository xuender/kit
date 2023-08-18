package cfg

import "errors"

var (
	ErrUnpaddingLength = errors.New("invalid unpadding length")
	ErrNoCipher        = errors.New("no cipher")
	ErrKey             = errors.New("password error")
)
