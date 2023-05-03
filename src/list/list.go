// Package list ...
package list

import (
	"fmt"
	"strings"
)

// List ...
type List[A any] struct {
	x    A
	xs   *List[A]
	size int
}

// FromSlice creates a new linked list with elements from a given slice.
// Time Complexity: O(n), where n is the length of the input slice.
// Space Complexity: O(n), where n is the length of the input slice.
// Example:
//
//	sl := []int{1, 2, 3}
//	lst := FromSlice(sl) // lst is a linked list with elements [1, 2, 3]
func FromSlice[A comparable](sl []A) *List[A] {
	var lst *List[A]
	for _, value := range sl {
		lst = lst.Append(value)
	}
	return lst
}

// ToSlice returns a slice containing the elements of the list in order.
// Time Complexity: O(n)
// Space Complexity: O(n)
// Example:
//
//	lst := FromSlice([]int{1, 2, 3})
//	sl := lst.ToSlice() // sl is []int{1, 2, 3}
func (lst *List[A]) ToSlice() []A {
	var result []A
	current := lst
	for current != nil {
		result = append(result, current.x)
		current = current.xs
	}
	return result
}

// Concat concatenates two lists and returns the resulting list.
// If either list is nil or empty, it returns the other list.
// The original lists are not modified.
//
// Time complexity: O(n), where n is the size of the first list.
// Space complexity: O(n), where n is the size of the first list.
//
// Example:
//
//	lst1 := FromSlice([]int{1, 2, 3}) // creates a list [1, 2, 3]
//	lst2 := FromSlice([]int{4, 5, 6}) // creates a list [4, 5, 6]
//	lst3 := lst1.Concat(lst2)         // concatenates lst1 and lst2 into lst3
//	result := lst3.ToSlice()          // returns [1, 2, 3, 4, 5, 6]
//
// TODO optimize to LogN
func (lst *List[A]) Concat(other *List[A]) *List[A] {
	if lst == nil || lst.size == 0 {
		return other
	}
	if other == nil || other.size == 0 {
		return lst
	}
	return &List[A]{x: lst.x, xs: lst.xs.Concat(other), size: lst.size + other.size}
}

// Append adds a new element to the end of the list and returns a new list.
// If the list is empty (nil), a new list with the given value is returned.
// Time Complexity: O(n), where n is the number of elements in the list.
// Space Complexity: O(n), where n is the number of elements in the list.
//
// Example:
//
//	l := &List{1, &List{2, nil}, 2}
//	l.Append(3) // &List{1, &List{2, &List{3, nil}}, 3}
func (lst *List[A]) Append(value A) *List[A] {
	if lst == nil {
		return &List[A]{x: value, size: 1}
	}
	return &List[A]{x: lst.x, xs: lst.xs.Append(value), size: lst.size + 1}
}

// AppendList appends the elements of the given list to the end of this list.
// If either this list or the given list is nil or empty, the other list is returned.
// The original lists are not modified; instead, a new list is returned.
//
// Time Complexity: O(n+m), where n and m are the sizes of the two lists.
// Space Complexity: O(n+m), where n and m are the sizes of the two lists.
//
// Example:
//
//	a := FromSlice([]int{1, 2, 3})
//	b := FromSlice([]int{4, 5, 6})
//	c := a.AppendList(b)
//	// a is still [1, 2, 3]
//	// b is still [4, 5, 6]
//	// c is [1, 2, 3, 4, 5, 6]
func (lst *List[A]) AppendList(toAppend *List[A]) *List[A] {
	// If this list is empty or nil, return the other list
	if lst == nil || lst.size == 0 {
		return toAppend
	}
	// If the list to append is empty or nil, return this list
	if toAppend == nil || toAppend.size == 0 {
		return lst
	}
	// Helper function that takes the current list and the list to append as arguments
	var appendHelper func(*List[A], *List[A]) *List[A]
	appendHelper = func(l *List[A], tail *List[A]) *List[A] {
		// If the current list is empty or nil, return the tail
		if l == nil || l.size == 0 {
			return tail
		}
		// Recursively append the tail to the rest of the list
		return &List[A]{x: l.x, xs: appendHelper(l.xs, tail), size: l.size + tail.size}
	}
	return appendHelper(lst, toAppend)
}

// Prepend adds a new element to the beginning of the List and returns a new List
// with the new element as the head and the original List as the tail.
// The time complexity of this operation is O(1) since it only involves creating a new List node
// with the new element and pointing it to the original List.
// The new List has a size that is one greater than the original List.
func (lst *List[A]) Prepend(value A) *List[A] {
	return &List[A]{x: value, xs: lst}
}

// PrependList prepends a given list to the beginning of the current list.
// The function returns a new list that represents the concatenation of the two lists.
// If either list is nil or has size 0, the other list is returned.
// Time complexity: O(logN)
// Space complexity: O(logN)
//
// Example:
// lst := &List[int]{x: 1, xs: &List[int]{x: 2, xs: &List[int]{x: 3, xs: nil}}, size: 3}
// toPrepend := &List[int]{x: 4, xs: &List[int]{x: 5, xs: nil}, size: 2}
// result := lst.PrependList(toPrepend)
// // result should be [4, 5, 1, 2, 3]
func (lst *List[A]) PrependList(toPrepend *List[A]) *List[A] {
	if toPrepend == nil || toPrepend.size == 0 {
		return lst
	}
	if lst == nil || lst.size == 0 {
		return toPrepend
	}
	// Split the first list into two lists at index 0
	left, right := lst.Split(0)
	// Concatenate the second list to the left side of the first list
	return toPrepend.Concat(left).Concat(right)
}

// Split splits the list into two sub-lists.
// The first list contains the first n elements of the original list,
// and the second list contains the remaining elements. The original list is not modified.
//
// Time complexity: O(n)
// Space complexity: O(n)
//
// Example:
//
// l := &List[int]{1, &List[int]{2, &List[int]{3, nil, 3}, 2}, 1}
// left, right := l.Split(1)
// fmt.Println(left)  // Output: [1]
// fmt.Println(right) // Output: [2 -> 3 -> nil]
func (lst *List[A]) Split(n int) (*List[A], *List[A]) {
	// Initialize left and right lists to nil
	var left, right *List[A]

	// Define a helper function to recursively split the list
	var helper func(*List[A], int) (*List[A], *List[A])
	helper = func(l *List[A], i int) (*List[A], *List[A]) {
		// Base case: if the list is empty or has size 0, return empty lists
		if l == nil || l.size == 0 {
			return &List[A]{size: 0}, &List[A]{size: 0}
		}
		// Base case: if n is 0, return the first element as left and the rest as right
		if i == 0 {
			return &List[A]{x: l.x, size: 1}, l.xs
		}
		// Recursive case: split the rest of the list (l.xs) and assign the first element (l.x) to the left list
		left1, right1 := helper(l.xs, i-1)
		return &List[A]{x: l.x, xs: left1, size: left1.size + 1}, right1
	}

	left, right = helper(lst, n-1)
	return left, right
}

// Reverse reverses the order of the elements in the list in-place.
// Time Complexity: O(n), where n is the number of elements in the list.
// Space Complexity: O(n), since the function creates a new list to store the reversed elements.
//
// Example:
// lst := &List[int]{1, 2, 3, 4, 5}
// reversed := lst.Reverse()
// fmt.Println(reversed) // Output: [5, 4, 3, 2, 1]
func (lst *List[A]) Reverse() *List[A] {
	reversed := &List[A]{}
	return lst.reverse(reversed)
}

func (lst *List[A]) reverse(acc *List[A]) *List[A] {
	if lst == nil {
		return acc
	}
	return lst.xs.reverse(&List[A]{x: lst.x, xs: acc, size: lst.size})
}

// At returns the element at the given index in the list, as well as a bool indicating whether the element was found.
// If the index is out of bounds or the list is empty, it returns the zero value for type A and false.
// The time complexity of this function is O(n), where n is the size of the list.
// The space complexity is O(1).
//
// Example usage:
// lst := &List[int]{1, 2, 3, 4, 5}
// x, ok := lst.At(2) // x is 3, ok is true
// y, ok := lst.At(5) // y is 0 (zero value of int), ok is false
func (lst *List[A]) At(index int) (A, bool) {
	var zero A
	if lst == nil || index < 0 || index >= lst.size {
		return zero, false
	}

	var helper func(*List[A], int) (A, bool)
	helper = func(l *List[A], i int) (A, bool) {
		if i == 0 {
			return l.x, true
		}
		return helper(l.xs, i-1)
	}
	return helper(lst, index)
}

// Head return the head element in the List
func (lst *List[A]) Head() A {
	return lst.x
}

// Tail return the tail of the List
func (lst *List[A]) Tail() *List[A] {
	return lst.xs
}

// Size return the size of the List
func (lst *List[A]) Size() int {
	return lst.size
}

// String give string representation of the List
func (lst *List[A]) String() string {
	if lst == nil || lst.size == 0 {
		return "[]"
	}
	var b strings.Builder
	b.WriteRune('[')
	b.WriteString(fmt.Sprintf("%v", lst.x))
	for cur := lst.xs; cur != nil; cur = cur.xs {
		b.WriteString(", ")
		b.WriteString(fmt.Sprintf("%v", cur.x))
	}
	b.WriteRune(']')
	return b.String()
}

// Filter returns a new list with the elements from the input list lst
// that satisfy the predicate fn.
//
// Time complexity: O(n), where n is the length of the input list lst.
// Filter applies the predicate fn to each element of lst exactly once.
//
// Space complexity: O(m), where m is the number of elements in the output list.
// Filter creates a new list with at most m elements, where m <= n.
//
// Example:
//
// lst := &List[int]{x: 1, xs: &List[int]{x: 2, xs: &List[int]{x: 3, xs: nil}}}
// fn := func(x int) bool { return x % 2 == 0 }
// result := lst.Filter(fn)
// // result is &List[int]{x: 2, xs: nil}
func (lst *List[A]) Filter(fn func(value A) bool) *List[A] {
	var acc *List[A]
	if ok := fn(lst.Head()); ok {
		acc = acc.Append(lst.Head())
	}
	p := lst
	for p.Tail() != nil {
		if ok := fn(p.Tail().Head()); ok {
			acc = acc.Append(p.Tail().Head())
		}
		p = p.Tail()
	}
	return acc
}

// Map applies a function fn to each element of the input list lst
// and returns a new list with the transformed elements.
// The input list lst is not modified.
//
// Time complexity: O(n), where n is the length of the input list lst.
// Map applies fn to each element of lst exactly once.
//
// Space complexity: O(n), where n is the length of the input list lst.
// Map creates a new list with the transformed elements, which has the same length as lst.
//
// Example:
//
// lst := &List[int]{x: 1, xs: &List[int]{x: 2, xs: &List[int]{x: 3, xs: nil}}}
// fn := func(x int) int { return x * x }
// result := Map(lst, fn)
// // result is &List[int]{x: 1, xs: &List[int]{x: 4, xs: &List[int]{x: 9, xs: nil}}}
func Map[A, B any](lst *List[A], fn func(value A) B) *List[B] {
	acc := &List[B]{x: fn(lst.Head()), size: 1}
	p := lst
	for p.Tail() != nil {
		acc = acc.Append(fn(p.Tail().Head()))
		p = p.Tail()
	}
	return acc
}
