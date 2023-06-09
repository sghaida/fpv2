// Package list ...
package list

import "github.com/sghaida/fpv2/src/collections"

// List is a type alias for a slice of elements of type A.
// The `A` type parameter specifies the type of elements that the slice can hold.
type List[A any] []A

// Ops is an interface that defines a set of operations that can be performed on a list of elements of type A.
// The `A` type parameter specifies the type of elements that the list can hold.
type Ops[A any] interface {
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
//
// Example usage:
// Given a list of integers, we can use the `Map` function to square each integer and create a new list of the squared values:
//
//	input := List[int]{1, 2, 3, 4, 5}
//	mapper := func(elm int) int { return elm * elm }
//	output := Map(input, mapper)
//	// Output: List[int]{1, 4, 9, 16, 25}
func Map[A, B any](lst List[A], mapper collections.Mapper[A, B]) List[B] {
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
//
// Example usage:
// Given a list of integers, we can use the `Reduce` function to compute the sum of the integers:
//
//	input := List[int]{1, 2, 3, 4, 5}
//	reducer := func(acc int, elm int) int { return acc + elm }
//	output := Reduce(input, reducer)
//	// Output: 15
func Reduce[A, B any](lst List[A], reducer collections.Reducer[A, B]) B {
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
//
// Example usage:
// Given a list of integers, we can use the `FoldLeft` function to compute the product of the integers:
//
//	input := List[int]{1, 2, 3, 4, 5}
//	reducer := func(acc int, elm int) int { return acc * elm }
//	output := FoldLeft(input, 1, reducer)
//	// Output: 120
func FoldLeft[A, B any](lst List[A], acc B, fn collections.Reducer[A, B]) B {
	for _, value := range lst {
		acc = fn(acc, value)
	}
	return acc
}

// Flatten takes a list of lists and returns a flattened list that contains all the elements in the original list of lists.
// The `A` type parameter specifies the type of elements in the original list of lists, which should be a list type `List[B]`.
// The returned list will have elements of type `B`.
//
// Example usage:
// Given a list of lists of integers, we can use the `Flatten` function to obtain a single list of all the integers:
//
//	input := List[List[int]]{
//	    {1, 2},
//	    {3},
//	    {4, 5, 6},
//	}
//	output := Flatten[int, int](input)
//	// Output: List[int]{1, 2, 3, 4, 5, 6}
func Flatten[A List[B], B any](lst List[A]) List[B] {
	var acc []B
	for _, sublist := range lst {
		for _, elem := range sublist {
			acc = append(acc, elem)
		}
	}
	return acc
}

// FlatMap returns a new list obtained by applying the given mapper function to each element of the input list
// and flattening the results. The mapper function takes an element of type A and returns a list of elements of
// type B. The resulting list is a concatenation of all the lists returned by the mapper function.
//
// Example usage:
//
//	lst := List{1, 2, 3}
//	mapper := func(x int) List[int] {
//	    return List{x, x * 2, x * 3}
//	}
//	result := FlatMap(lst, mapper) // result = List{1, 2, 3, 2, 4, 6, 3, 6, 9}
//
// The FlatMap function is implemented in terms of the Map and Flatten functions.
// See their respective documentation for more information.
func FlatMap[A, B any](lst List[A], mapper collections.Mapper[A, List[B]]) List[B] {
	return Flatten(Map(lst, mapper))
}

// Filter returns a new list containing only the elements of the input list that satisfy the given predicate function.
// The predicate function takes an element of type A and returns true if the element should be included in the resulting list,
// false otherwise.
//
// Example usage:
//
//	lst := List{1, 2, 3, 4, 5}
//	predicate := func(x int) bool {
//	    return x % 2 == 0
//	}
//	result := Filter(lst, predicate) // result = List{2, 4}
//
// The Filter function is implemented using a loop that iterates over each element of the input list and checks whether
// the predicate function returns true for that element. If it does, the element is added to a new list that is returned
// as the result of the function.
func Filter[A any](lst List[A], predicate func(A) bool) List[A] {
	var acc []A
	for _, elm := range lst {
		if predicate(elm) {
			acc = append(acc, elm)
		}
	}
	return acc
}
