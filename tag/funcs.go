package tag

import "golang.org/x/exp/constraints"

// Get 获取标签.
func Get[E constraints.Integer](tag E) []int {
	ret := []int{}

	for idx := 0; ; idx++ {
		val := E(1 << idx)
		if val > tag {
			return ret
		}

		if tag&val > 0 {
			ret = append(ret, idx)
		}
	}
}

func GetBit[E constraints.Integer](tag E) []E {
	ret := []E{}

	for idx := E(0); ; idx++ {
		val := E(1 << idx)
		if val > tag {
			return ret
		}

		if tag&val > 0 {
			ret = append(ret, val)
		}
	}
}

// Tag 合并成标签.
func Tag[E constraints.Integer](elems ...int) E {
	return Add(E(0), elems...)
}

// Add 增加.
func Add[E constraints.Integer](tag E, elems ...int) E {
	for _, elem := range elems {
		// OR
		tag |= (1 << elem)
	}

	return tag
}

// Del 删除.
func Del[E constraints.Integer](tag E, elems ...int) E {
	for _, elem := range elems {
		// AND NOT
		tag &^= (1 << elem)
	}

	return tag
}

// Has 包含任何一个标签.
func Has[E constraints.Integer](tag E, elems ...int) bool {
	for _, elem := range elems {
		// AND
		if tag&(1<<elem) > 0 {
			return true
		}
	}

	return false
}

// Hit 命中所有标签.
func Hit[E constraints.Integer](tag E, elems ...int) bool {
	return !Has(^tag, elems...)
}
