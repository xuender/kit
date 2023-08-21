package los_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/los"
)

func TestSunday_IndexOf(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Equal(0, los.NewSunday([]rune("a")).IndexOf([]rune("a")))
	ass.Equal(1, los.NewSunday([]rune("abc")).IndexOf([]rune("babc")))
	ass.Equal(1, los.NewSunday([]rune("a")).IndexOf([]rune("ba")))
	ass.Equal(0, los.NewSunday([]rune("")).IndexOf([]rune("abcaabcab")))
	ass.Equal(-1, los.NewSunday([]rune("aa")).IndexOf([]rune("b")))
	ass.Equal(1, los.NewSunday([]rune("a")).IndexOf([]rune("bab")))
	ass.Equal(0, los.NewSunday([]rune("a")).IndexOf([]rune("ab")))
	ass.Equal(4, los.NewSunday([]rune("abcab")).IndexOf([]rune("abcaabcab")))
	ass.Equal(-1, los.NewSunday([]rune("1abcab")).IndexOf([]rune("abcaabcab")))
	ass.Equal(-1, los.NewSunday([]rune("1")).IndexOf([]rune("abcaabcab")))
}
