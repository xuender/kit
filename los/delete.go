package los

import "slices"

func Delete[S ~[]E, E any](slice S, idxs ...int) S {
	switch len(idxs) {
	case 0:
		return slice
	case 1:
		return slices.Delete(slice, idxs[0], idxs[0]+1)
	default:
		slices.Sort(idxs)
		slices.Reverse(idxs)

		end := len(slice)
		start := len(slice)

		for _, idx := range idxs {
			if idx == start-1 {
				start = idx

				continue
			}

			if start < end {
				slice = slices.Delete(slice, start, end)
			}

			end = idx + 1
			start = idx
		}

		return slices.Delete(slice, start, end)
	}
}

func DeleteBy[S ~[]E, E comparable](slice S, elems ...E) S {
	switch len(elems) {
	case 0:
		return slice
	case 1:
		return Delete(slice, IndexAll(slice, elems[0])...)
	default:
		return DeleteFunc(slice, func(elem E) bool { return slices.Contains(elems, elem) })
	}
}

func DeleteFunc[S ~[]E, E any](slice S, del func(E) bool) S {
	return Delete(slice, IndexAllFunc(slice, del)...)
}
