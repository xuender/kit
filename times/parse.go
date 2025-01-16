package times

import (
	"errors"
	"time"
)

var ErrFormat = errors.New("time format error")

// NewParse returns a function that attempts to parse a string into a time.Time.
func NewParse() func(string) (time.Time, error) {
	layouts := map[int][]string{}
	defaultLayout := time.DateTime

	for _, layout := range []string{
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,
		"0102",
		"2006",
		"20060102",
		"060102",
		"6-01-02",
		"06-01-02",
		"060102150405",
		"060102 150405",
		"20060102150405",
		"20060102 150405",
		"06-01-02 15:04:05",
		"2006-01-02T15:04:05",
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	} {
		length := len(layout)
		layouts[length] = append(layouts[length], layout)
	}

	return func(str string) (time.Time, error) {
		length := len(str)
		if length == len(defaultLayout) {
			if val, err := time.Parse(defaultLayout, str); err == nil {
				return val, nil
			}
		}

		if items, has := layouts[length]; has {
			for _, layout := range items {
				if val, err := time.Parse(layout, str); err == nil {
					defaultLayout = layout

					return val, nil
				}
			}
		}

		return time.Now(), ErrFormat
	}
}
