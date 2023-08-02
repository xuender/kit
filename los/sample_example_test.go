package los_test

import (
	"fmt"
	"math/rand"

	"github.com/xuender/kit/los"
)

func ExampleSampleBool() {
	rand.Seed(3)
	fmt.Println(los.SampleBool())
	fmt.Println(los.SampleBool())

	// Output:
	// false
	// true
}
