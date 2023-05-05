package collections

// Reducer is a type that represents a function that reduces a list of values of type A to a single value of type B.
// The `acc` parameter represents the accumulated value, and `value` represents the current value being processed.
// The function returns the updated accumulated value after processing the current value.
// The types of `A` and `B` are specified using Go's `any` keyword, allowing for any type to be used.
type Reducer[A, B any] func(acc B, value A) B

// Mapper is a type that represents a function that maps an input value of type A to an output value of type B.
// The `elm` parameter represents the input value to be mapped, and the function returns the corresponding output value.
// The types of `A` and `B` are specified using Go's `any` keyword, allowing for any type to be used.
type Mapper[A, B any] func(elm A) B
