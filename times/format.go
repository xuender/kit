package times

import "time"

const (
	_date   = "2006-01-02"
	_format = "2006-01-02 15:04:05"
	_time   = "15:04:05"
)

func Format(input time.Time) string {
	return input.Format(_format)
}

func FormatDate(input time.Time) string {
	return input.Format(_date)
}

func FormatTime(input time.Time) string {
	return input.Format(_time)
}
