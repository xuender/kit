package los

func Map[S ~[]T, T any, R any](slice S, iteratee func(item T) R) []R {
	result := make([]R, len(slice))

	for i, item := range slice {
		result[i] = iteratee(item)
	}

	return result
}
