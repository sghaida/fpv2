// Package list ...
package list

type List[A comparable] struct {
	x    A
	xs   *List[A]
	size int
}

// FromSlice create a list from slice
func FromSlice[A comparable](sl []A) *List[A] {
	var lst *List[A]
	for _, value := range sl {
		lst = lst.Append(value)
	}
	return lst
}

func (lst *List[A]) ToSlice() []A {
	var result []A
	current := lst
	for current != nil {
		result = append(result, current.x)
		current = current.xs
	}
	return result
}

// Append appends element to the List
func (lst *List[A]) Append(value A) *List[A] {
	if lst == nil {
		return &List[A]{x: value, size: 1}
	}
	return &List[A]{x: lst.x, xs: lst.xs.Append(value), size: lst.size + 1}
}

// Prepend prepends an element to the List
func (lst *List[A]) Prepend(value A) *List[A] {
	return &List[A]{x: value, xs: lst}
}

// Split splits the List on n to 2 Lists
func (lst *List[A]) Split(n int) (*List[A], *List[A]) {
	if lst == nil {
		return &List[A]{}, &List[A]{}
	}
	var left, right *List[A]
	right = lst
	for i := 0; i < n && right != nil && right.size != 0; i++ {
		left = left.Append(right.x)
		right = right.xs
	}
	if left == nil {
		left = &List[A]{}
	}
	return left, right
}

// Reverse reverses the order of the List O(N)
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

// At return the List at specific index O(Log(N))
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
