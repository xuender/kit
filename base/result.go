package base

// Result1 返回第1个参数.
func Result1[R any](r R, _ ...any) R {
	return r
}

// Result2 返回第2个参数.
func Result2[R any](_ any, r R, _ ...any) R {
	return r
}

// Result3 返回第3个参数.
func Result3[R any](_, _ any, r R, _ ...any) R {
	return r
}
