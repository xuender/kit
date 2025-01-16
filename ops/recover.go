package ops

import (
	"errors"
	"fmt"
)

// Recover is a panic recovery function that wraps the execution of a given function.
// It captures any panics that occur during the execution of the function and returns them as errors.
func Recover(yield func(error)) {
	// nolint: err113
	if msg := recover(); msg != nil {
		switch val := msg.(type) {
		case error:
			yield(val)
		case string:
			yield(errors.New(val))
		default:
			yield(fmt.Errorf("%v", msg))
		}
	}
}
