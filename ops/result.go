package ops

// Result returns the provided argument unchanged.
func Result[R any](arg R, _ ...any) R {
	return arg
}

// Result1 returns the provided argument unchanged.
func Result1[R any](arg R, _ ...any) R {
	return arg
}

// Result2 returns the second provided argument unchanged.
func Result2[R any](_ any, arg R, _ ...any) R {
	return arg
}

// Result3 returns the third provided argument unchanged.
func Result3[R any](_, _ any, arg R, _ ...any) R {
	return arg
}

// Result4 returns the fourth provided argument unchanged.
func Result4[R any](_, _, _ any, arg R, _ ...any) R {
	return arg
}
