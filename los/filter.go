package los

func Filter[S ~[]V, V any](slice S, predicate func(item V) bool) S {
	idxs := IndexAllFunc(slice, predicate)
	if len(idxs) == 0 {
		return nil
	}

	result := make([]V, len(idxs))
	for i, idx := range idxs {
		result[i] = slice[idx]
	}

	return result
}
