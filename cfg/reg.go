package cfg

import "regexp"

var (
	_regSecret    = regexp.MustCompile(`(?:A|D)ES\((?:\w|/|\+)+=*\)`)
	_regPlaintext = regexp.MustCompile(`(?:A|D)ES\[(?:\w|/|\+)+=*\]`)
	_regKey       = regexp.MustCompile(`(?:A|D)ES(?:\(|\[)(?:\w|/|\+)+=*(?:\)|\])`)
)
