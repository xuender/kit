package los

import (
	"math/rand"

	"github.com/xuender/kit/base"
)

func SampleBool() bool {
	// nolint: gosec
	return rand.Intn(base.Two) > 0
}
