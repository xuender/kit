package base

// Only21 返回2个返回值第1个.
func Only21[R, T any](r R, _ T) R {
	return r
}

// Only22 返回2个返回值第2个.
func Only22[R, T any](_ T, r R) R {
	return r
}

// Only31 返回3个返回值第1个.
func Only31[R, T1, T2 any](r R, _ T1, _ T2) R {
	return r
}

// Only32 返回3个返回值第2个.
func Only32[R, T1, T2 any](_ T1, r R, _ T2) R {
	return r
}

// Only33 返回3个返回值第3个.
func Only33[R, T1, T2 any](_ T1, _ T2, r R) R {
	return r
}
