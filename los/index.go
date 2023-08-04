package los

func IndexOf[T comparable](slice, sub []T) int {
	if len(slice) < len(sub) {
		return -1
	}

	for index := 0; index <= len(slice)-len(sub); index++ {
		has := true

		for elemIndex, elem := range sub {
			if slice[index+elemIndex] != elem {
				has = false

				break
			}
		}

		if has {
			return index
		}
	}

	return -1
}

func LastIndexOf[T comparable](slice, sub []T) int {
	lenSub := len(sub)
	if len(slice) < lenSub {
		return -1
	}

	for index := len(slice) - lenSub; index >= 0; index-- {
		has := true

		for elemIndex, elem := range sub {
			if slice[index+elemIndex] != elem {
				has = false

				break
			}
		}

		if has {
			return index
		}
	}

	return -1
}
