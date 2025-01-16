package types

// Checker defines a function type that tests a value and returns a boolean.
type Checker[T any] func(T) bool
