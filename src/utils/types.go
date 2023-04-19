package utils

// Mapper is a function that applies on type A and transform it to type B
type Mapper[A, B any] func(A) B

// Predicate function used for filters and checking
type Predicate[A any] func(value A) bool
