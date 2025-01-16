package types

// Mapper defines a function type that transforms a value.
type Mapper[K, V any] func(K) V
