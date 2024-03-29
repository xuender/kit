package los

import (
	"fmt"
)

func Must[T any](val T, err error, messageArgs ...any) T {
	must(err, messageArgs...)

	return val
}

func Must0(err error, messageArgs ...any) {
	must(err, messageArgs...)
}

func Must1[T any](val T, err error, messageArgs ...any) T {
	return Must(val, err, messageArgs...)
}

func Must2[T1, T2 any](val1 T1, val2 T2, err error, messageArgs ...any) (T1, T2) {
	must(err, messageArgs...)

	return val1, val2
}

func Must3[T1, T2, T3 any](val1 T1, val2 T2, val3 T3, err error, messageArgs ...any) (T1, T2, T3) {
	must(err, messageArgs...)

	return val1, val2, val3
}

func Must4[T1, T2, T3, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, err error, messageArgs ...any) (T1, T2, T3, T4) {
	must(err, messageArgs...)

	return val1, val2, val3, val4
}

func Must5[T1, T2, T3, T4, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5,
	err error, messageArgs ...any,
) (T1, T2, T3, T4, T5) {
	must(err, messageArgs...)

	return val1, val2, val3, val4, val5
}

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
