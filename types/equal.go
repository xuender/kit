package types

// Equaler is a function type that defines how to compare two values of type E for equality.
type Equaler[E any] func(E, E) bool
