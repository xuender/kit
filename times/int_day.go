package times

import (
	"strconv"
	"time"
)

const (
	_intDayFormat = "20060102"
	_year         = 10000
	_month        = 100
)

type IntDay int

func Time2IntDay(input time.Time) IntDay {
	return IntDay(input.Year()*_year + int(input.Month())*_month + input.Day())
}

func Now2IntDay() IntDay {
	return Time2IntDay(time.Now())
}

func ParseIntDay(str string) (IntDay, error) {
	day, err := time.Parse(_intDayFormat, str)
	if err != nil {
		return 0, err
	}

	return Time2IntDay(day), nil
}

func (p IntDay) String() string {
	return strconv.Itoa(int(p))
}

func (p IntDay) Day() int {
	return int(p) % _month
}

func (p IntDay) Month() int {
	return int(p) % _year / _month
}

func (p IntDay) Year() int {
	return int(p) / _year
}
