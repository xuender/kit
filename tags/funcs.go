package tags

import "golang.org/x/exp/constraints"

// Add 增加.
func Add[E constraints.Integer](elems ...E) E {
	var tag E

	for _, elem := range elems {
		tag |= elem
	}

	return tag
}

// Del 删除.
func Del[E constraints.Integer](tag E, elems ...E) E {
	for _, elem := range elems {
		tag ^= (tag & elem)
	}

	return tag
}

// Hit 命中所有标签.
func Hit[E constraints.Integer](tag E, elems ...E) bool {
	for _, elem := range elems {
		if ^tag&elem > 0 {
			return false
		}
	}

	return true
}

// Has 包含任何一个标签.
func Has[E constraints.Integer](tag E, elems ...E) bool {
	for _, elem := range elems {
		if tag&elem > 0 {
			return true
		}
	}

	return false
}
