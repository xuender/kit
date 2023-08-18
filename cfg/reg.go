package cfg

import "regexp"

var (
	_regSecret    = regexp.MustCompile(`(A|D)ES\([\w\+/-]+=*\)`)
	_regPlaintext = regexp.MustCompile(`(A|D)ES\[[\p{Han}\w\x20-\x27\x2A-\x40]+\]`)
	_regCheck     = regexp.MustCompile(`^[\p{Han}\w\x20-\x27\x2A-\x40]+$`)
)
