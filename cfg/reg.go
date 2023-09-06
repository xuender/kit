package cfg

import "regexp"

var (
	_secretRegex    = regexp.MustCompile(`(A|D)ES\([\w\+/-]+=*\)`)
	_plaintextRegex = regexp.MustCompile(`(A|D)ES\[[\p{Han}\w\x20-\x27\x2A-\x40]+\]`)
	_checkRegex     = regexp.MustCompile(`^[\p{Han}\w\x20-\x27\x2A-\x40]+$`)
)
