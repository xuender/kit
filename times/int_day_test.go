package times_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/times"
)

func TestParseIntDay(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	_, err := times.ParseIntDay("error")

	ass.NotNil(err)
}

func TestIntDay_Marshal(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	day := times.IntDay(20230918)

	ass.Equal([]byte{0x01, 0x34, 0xB3, 0x06}, day.Marshal())

	day = times.IntDay(10230918)

	ass.Equal([]byte{0x00, 0x9C, 0x1C, 0x86}, day.Marshal())
}

func TestIntDay_MarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	day := times.IntDay(20230918)

	ass.Equal([]byte{0x32, 0x30, 0x32, 0x33, 0x30, 0x39, 0x31, 0x38}, lo.Must1(day.MarshalJSON()))

	day = times.IntDay(10230918)

	ass.Equal([]byte{0x31, 0x30, 0x32, 0x33, 0x30, 0x39, 0x31, 0x38}, lo.Must1(day.MarshalJSON()))
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	var day times.IntDay

	err := day.UnmarshalJSON([]byte{0x32, 0x30, 0x32, 0x33, 0x30, 0x39, 0x31, 0x38})
	ass.Nil(err)
	ass.Equal(20230918, int(day))

	err = day.UnmarshalJSON([]byte{0x31, 0x30, 0x32, 0x33, 0x30, 0x39, 0x31, 0x38})

	ass.Nil(err)
	ass.Equal(10230918, int(day))
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	var day times.IntDay

	err := day.Unmarshal([]byte{0x01, 0x34, 0xB3, 0x06})
	ass.Nil(err)
	ass.Equal(20230918, int(day))

	err = day.Unmarshal([]byte{0x00, 0x9C, 0x1C, 0x86})

	ass.Nil(err)
	ass.Equal(10230918, int(day))
}
