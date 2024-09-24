package times

import (
	"encoding/binary"
	"encoding/json"
	"strconv"
	"time"
)

const (
	_intDayFormat = "20060102"
	_year         = 10_000
	_month        = 100
)

type IntDay int32

func Time2IntDay(input time.Time) IntDay {
	// nolint: gosec
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

func (p IntDay) Marshal() []byte {
	const length = 4

	data := make([]byte, length)
	// nolint: gosec
	binary.BigEndian.PutUint32(data, uint32(p))

	return data
}

func (p IntDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(int32(p))
}

func (p *IntDay) UnmarshalJSON(data []byte) error {
	var day int32

	err := json.Unmarshal(data, &day)
	*p = IntDay(day)

	return err
}

func (p *IntDay) Unmarshal(data []byte) error {
	if length := 4; len(data) < length {
		return ErrParseError
	}
	// nolint: gosec
	*p = IntDay(binary.BigEndian.Uint32(data))

	return nil
}
