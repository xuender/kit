// nolint: gosec
package ordered_test

import (
	"math/rand"
	"testing"

	"github.com/xuender/kit/ordered"
	"github.com/xuender/kit/set"
)

func BenchmarkMap(b *testing.B) {
	set := set.NewSet[int]()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		set.Add(rand.Int())
	}
}

func BenchmarkSet(b *testing.B) {
	set := ordered.NewSet[int]()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		set.Add(rand.Int())
	}
}
