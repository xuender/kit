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
