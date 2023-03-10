package base

// EveryNil 是否全是 nil.
func EveryNil(elems ...any) bool {
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

// SomeNil 是否包含 nil.
func SomeNil(elems ...any) bool {
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
