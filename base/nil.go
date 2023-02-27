package base

// AllNil 是否全是 nil.
func AllNil(elems ...any) bool {
	if len(elems) == 0 {
		return true
	}

	for _, elem := range elems {
		if elem != nil {
			return false
		}
	}

	return true
}

// AnyNil 是否包含 nil.
func AnyNil(elems ...any) bool {
	if len(elems) == 0 {
		return true
	}

	for _, elem := range elems {
		if elem == nil {
			return true
		}
	}

	return false
}
