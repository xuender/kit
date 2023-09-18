package times_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/times"
)

func TestParseIntDay(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	_, err := times.ParseIntDay("error")

	ass.NotNil(err)
}

func TestIntDay_Bytes(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	day := times.IntDay(20230918)

	ass.Equal([]byte{0x01, 0x34, 0xB3, 0x06}, day.Bytes())

	day = times.IntDay(10230918)

	ass.Equal([]byte{0x00, 0x9C, 0x1C, 0x86}, day.Bytes())
}
