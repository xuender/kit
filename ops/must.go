package ops

import (
	"fmt"
)

// Must ensures that an error is nil. If not, it panics with the provided message arguments.
// It returns the original value.
func Must[T any](val T, err error, messageArgs ...any) T {
	must(err, messageArgs...)

	return val
}

// Must0 ensures that an error is nil and panics if it is not, using the provided message arguments.
func Must0(err error, messageArgs ...any) {
	must(err, messageArgs...)
}

// Must1 ensures that an error is nil and returns the original value. If the error is not nil, it panics.
func Must1[T any](val T, err error, messageArgs ...any) T {
	return Must(val, err, messageArgs...)
}

// Must2 ensures that an error is nil and returns two values. If the error is not nil, it panics.
func Must2[T1, T2 any](val1 T1, val2 T2, err error, messageArgs ...any) (T1, T2) {
	must(err, messageArgs...)

	return val1, val2
}

// Must3 ensures that an error is nil and returns three values. If the error is not nil, it panics.
func Must3[T1, T2, T3 any](val1 T1, val2 T2, val3 T3, err error, messageArgs ...any) (T1, T2, T3) {
	must(err, messageArgs...)

	return val1, val2, val3
}

// Must4 ensures that an error is nil and returns four values. If the error is not nil, it panics.
func Must4[T1, T2, T3, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, err error, messageArgs ...any) (T1, T2, T3, T4) {
	must(err, messageArgs...)

	return val1, val2, val3, val4
}

// Must5 ensures that an error is nil and returns five values. If the error is not nil, it panics.

func Must5[T1, T2, T3, T4, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5,
	err error, messageArgs ...any,
) (T1, T2, T3, T4, T5) {
	must(err, messageArgs...)

	return val1, val2, val3, val4, val5
}

// Must6 ensures that an error is nil and returns six values. If the error is not nil, it panics.
func Must6[T1, T2, T3, T4, T5, T6 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, val6 T6,
	err error, messageArgs ...any,
) (T1, T2, T3, T4, T5, T6) {
	must(err, messageArgs...)

	return val1, val2, val3, val4, val5, val6
}

func must(err error, messageArgs ...any) {
	if err == nil {
		return
	}

	if len(messageArgs) > 0 {
		if format, ok := messageArgs[0].(string); ok {
			panic(fmt.Sprintf(format, messageArgs[1:]...))
		}
	}

	panic(err)
}
