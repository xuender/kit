package los_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/xuender/kit/los"
)

func BenchmarkIndexOf(b *testing.B) {
	sub := lo.RepeatBy(100, func(_ int) rune { return rune('1') })
	sub[0] = 'a'
	slice := append(lo.RepeatBy(1_000_000, func(_ int) rune { return rune('0') }), sub...)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		los.IndexOf(slice, sub)
	}
}

func BenchmarkSunday_IndexOf(b *testing.B) {
	sub := lo.RepeatBy(100, func(_ int) rune { return rune('1') })
	sub[0] = 'a'
	slice := append(lo.RepeatBy(1_000_000, func(_ int) rune { return rune('0') }), sub...)
	sunday := los.NewSunday(sub)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sunday.IndexOf(slice)
	}
}
