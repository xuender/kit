package tags

import "golang.org/x/exp/constraints"

// Add 增加.
func Add[E constraints.Integer](tag E, elems ...E) E {
	for _, elem := range elems {
		// OR
		tag |= elem
	}

	return tag
}

// Del 删除.
func Del[E constraints.Integer](tag E, elems ...E) E {
	for _, elem := range elems {
		// AND NOT
		tag &^= elem
	}

	return tag
}

// Has 包含任何一个标签.
func Has[E constraints.Integer](tag E, elems ...E) bool {
	for _, elem := range elems {
		// AND
		if tag&elem > 0 {
			return true
		}
	}

	return false
}

// Hit 命中所有标签.
func Hit[E constraints.Integer](tag E, elems ...E) bool {
	return !Has(^tag, elems...)
}
