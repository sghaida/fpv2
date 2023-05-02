// Package collections ...
package collections

// List is a type alias for a slice of elements of type A.
// The `A` type parameter specifies the type of elements that the slice can hold.
type List[A any] []A

// Reducer is a type that represents a function that reduces a list of values of type A to a single value of type B.
// The `acc` parameter represents the accumulated value, and `value` represents the current value being processed.
// The function returns the updated accumulated value after processing the current value.
// The types of `A` and `B` are specified using Go's `any` keyword, allowing for any type to be used.
type Reducer[A, B any] func(acc B, value A) B

// Mapper is a type that represents a function that maps an input value of type A to an output value of type B.
// The `elm` parameter represents the input value to be mapped, and the function returns the corresponding output value.
// The types of `A` and `B` are specified using Go's `any` keyword, allowing for any type to be used.
type Mapper[A, B any] func(elm A) B

// ListOps is an interface that defines a set of operations that can be performed on a list of elements of type A.
// The `A` type parameter specifies the type of elements that the list can hold.
type ListOps[A any] interface {
	// Size returns the number of elements in the list.
	Size() int
	// Foreach applies the given function `fn` to each element in the list.
	Foreach(fn func(elm A))
	// Append returns a new list that is the result of appending the given element `elm` to the end of the current list.
	Append(elm A) List[A]
	// Prepend returns a new list that is the result of prepending the given element `elm` to the beginning of the current list.
	Prepend(elm A) List[A]
	// Take returns a new list that contains the first `n` elements of the current list.
	Take(n int) List[A]
}

// Size returns the number of elements in the list.
func (a List[A]) Size() int {
	return len(a)
}

// Foreach applies the given function `fn` to each element in the list.
func (a List[A]) Foreach(fn func(elm A)) {
	for _, value := range a {
		fn(value)
	}
}

// Append returns a new list that is the result of appending the given element `elm` to the end of the current list.
func (a List[A]) Append(elm A) List[A] {
	arr := make(List[A], len(a), len(a)+1)
	copy(arr, a)
	arr = append(arr, elm)
	return arr
}

// Prepend returns a new list that is the result of prepending the given element `elm` to the beginning of the current list.
func (a List[A]) Prepend(elm A) List[A] {
	arr := make(List[A], 0, len(a)+1)
	arr = append(arr, elm)
	arr = append(arr, a...)
	return arr
}

// Take returns a new list that contains the first `n` elements of the current list.
func (a List[A]) Take(n int) List[A] {
	if n < 0 {
		return make(List[A], 0)
	}
	arr := make(List[A], 0, n)
	for i := 0; i < n && i < a.Size(); i++ {
		arr = append(arr, a[i])
	}
	return arr
}

// Map applies the given `mapper` function to each element in the `lst` list, and returns a new list
// containing the mapped elements. The `A` type parameter specifies the type of elements in the original list,
// while the `B` type parameter specifies the type of elements in the mapped list.
// The `mapper` function takes an element of type `A` as input, and returns an element of type `B` as output.
// The function returns a new list containing the mapped elements.
// Example usage:
// Given a list of integers, we can use the `Map` function to square each integer and create a new list of the squared values:
//
//	input := List[int]{1, 2, 3, 4, 5}
//	mapper := func(elm int) int { return elm * elm }
//	output := Map(input, mapper)
//	// Output: List[int]{1, 4, 9, 16, 25}
func Map[A, B any](lst List[A], mapper Mapper[A, B]) List[B] {
	acc := make([]B, 0, len(lst))
	for _, elm := range lst {
		acc = append(acc, mapper(elm))
	}
	return acc
}

// Reduce applies the given `reducer` function to each element in the `lst` list and returns a single value of type `B`.
// The `A` type parameter specifies the type of elements in the original list,
// while the `B` type parameter specifies the type of the accumulated value.
// The `reducer` function takes an accumulated value of type `B` and an element of type `A` as input,
// and returns a new accumulated value of type `B` as output.
// The function applies the `reducer` function to each element in the `lst` list and returns the final accumulated value.
// Example usage:
// Given a list of integers, we can use the `Reduce` function to compute the sum of the integers:
//
//	input := List[int]{1, 2, 3, 4, 5}
//	reducer := func(acc int, elm int) int { return acc + elm }
//	output := Reduce(input, reducer)
//	// Output: 15
func Reduce[A, B any](lst List[A], reducer Reducer[A, B]) B {
	var acc B
	return FoldLeft(lst, acc, reducer)
}

// FoldLeft applies the given `fn` reducer function to each element in the `lst` list, starting with the initial accumulator value `acc`,
// and returns the final accumulated value of type `B`.
// The `A` type parameter specifies the type of elements in the original list,
// while the `B` type parameter specifies the type of the accumulated value.
// The `fn` reducer function takes an accumulated value of type `B` and an element of type `A` as input,
// and returns a new accumulated value of type `B` as output.
// The function applies the `fn` reducer function to each element in the `lst` list,
// starting with the initial accumulator value `acc`, and returns the final accumulated value.
// Example usage:
// Given a list of integers, we can use the `FoldLeft` function to compute the product of the integers:
//
//	input := List[int]{1, 2, 3, 4, 5}
//	reducer := func(acc int, elm int) int { return acc * elm }
//	output := FoldLeft(input, 1, reducer)
//	// Output: 120
func FoldLeft[A, B any](lst List[A], acc B, fn Reducer[A, B]) B {
	for _, value := range lst {
		acc = fn(acc, value)
	}
	return acc
}
