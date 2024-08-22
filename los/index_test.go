package los_test

import (
	"slices"
	"testing"

	"github.com/xuender/kit/los"
)

func BenchmarkIndexOf(b *testing.B) {
	sub := slices.Repeat([]rune{'1'}, 100)
	sub[0] = 'a'
	slice := append(slices.Repeat([]rune{'0'}, 1_000_000), sub...)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		los.IndexOf(slice, sub)
	}
}

func BenchmarkSunday_IndexOf(b *testing.B) {
	sub := slices.Repeat([]rune{'1'}, 100)
	sub[0] = 'a'
	slice := append(slices.Repeat([]rune{'0'}, 1_000_000), sub...)
	sunday := los.NewSunday(sub)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sunday.IndexOf(slice)
	}
}
