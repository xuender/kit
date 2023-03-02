package base

// Result1 返回第1个结果.
func Result1[R any](r R, _ ...any) R {
	return r
}

// Result2 返回第2个结果.
func Result2[R any](_ any, r R, _ ...any) R {
	return r
}

// Result3 返回第3个结果.
func Result3[R any](_, _ any, r R, _ ...any) R {
	return r
}

// Result4 返回第4个结果.
func Result4[R any](_, _, _ any, r R, _ ...any) R {
	return r
}
