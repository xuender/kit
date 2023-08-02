package times_test

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/times"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := [][]string{
		{"19780321", "20060102"},
		{"780321", "060102"},
		{"1978-03-21", "2006-01-02"},
		{"1978-03-21 10:21:34", "2006-01-02 15:04:05"},
	}

	for _, tt := range tests {
		day := lo.Must1(times.Parse(tt[0]))
		want := lo.Must1(time.Parse(tt[1], tt[0]))
		assert.Equal(t, day, want)
	}

	_, err := times.Parse("11")
	assert.NotNil(t, err)
}

func TestParse_Error(t *testing.T) {
	t.Parallel()

	_, err := times.Parse("23424fasdkfjk")
	assert.NotNil(t, err)
}

func TestParseNumber(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 2022, times.ParseNumber(1648777840.2672896).Year())
	assert.Equal(t, 2022, times.ParseNumber(1648777840000).Year())
	assert.Equal(t, 2022, times.ParseNumber(1648777840).Year())
	assert.Equal(t, 2022, times.ParseNumber(1648777).Year())
	assert.Equal(t, 2022, times.ParseNumber(20220401).Year())
	assert.Equal(t, 2022, times.ParseNumber(220401).Year())
	assert.Equal(t, 2002, times.ParseNumber(20401).Year())
	assert.Equal(t, time.Month(11), times.ParseNumber(1101).Month())
	assert.Equal(t, time.Month(4), times.ParseNumber(401).Month())
}
