package types

// Reducer defines a function type that reduces two values to one.
type Reducer[T any] func(T, T) T
