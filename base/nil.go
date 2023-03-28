package base

// EveryNil 检查所有元素是否都为 nil.
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

// SomeNil 检查是否有元素为 nil.
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
