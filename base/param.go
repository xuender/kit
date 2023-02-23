package base

// Param1 返回第1个参数.
func Param1[R any](r R, _ ...any) R {
	return r
}

// Param2 返回第2个参数.
func Param2[R any](_ any, r R, _ ...any) R {
	return r
}

// Param3 返回第3个参数.
func Param3[R any](_, _ any, r R, _ ...any) R {
	return r
}
