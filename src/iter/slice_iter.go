// Package iter ...
package iter

// SliceOps include the operations that can be done on a SliceIter
type SliceOps[A any] interface {
	Clone() SliceIter[A]
	Contains(elm A) bool
	Drop(n int) SliceIter[A]
	Filter(fn func(A) bool) SliceIter[A]
	Fold(fn func(A, A) A) A
	FoldLeft(initialValue any, fn func(any, A) any) any
	Foreach(fn func(A))
	Map(fn func(A) any) Iter[any]
	Reduce(fn func(A, A) A) A
	Take(n int) SliceIter[A]
	ToSlice() []A
	ToIter() Iter[A]
	Slice(from, until int) SliceIter[A]
}

// SliceIter definition of Slice Iterator
type SliceIter[A any] interface {
	Iter[A]
	SliceOps[A]
}

type sliceIter[A any] struct {
	slice   []A
	size    int
	current int
}

// FromSlice creates Iter from slice
func FromSlice[A any](slice []A) SliceIter[A] {
	return &sliceIter[A]{slice: slice, size: len(slice), current: 0}
}

// HasNext check if there is next element
func (si *sliceIter[A]) HasNext() bool {
	if len(si.slice) != 0 {
		return true
	}
	return false
}

// Next return the next element in the slice if available
func (si *sliceIter[A]) Next() A {
	if !si.HasNext() {
		var zero A
		return zero
	}
	item := si.slice[0]
	si.slice = si.slice[1:]
	si.current++
	si.size = len(si.slice)
	return item
}

// Count return the size of the iter and move to the end of the iter
func (si *sliceIter[A]) Count() int {
	count := len(si.slice)
	si.slice = []A{}
	si.current = si.size
	return count
}

// Size return the size of the iter
func (si *sliceIter[A]) Size() int {
	return si.size
}

// ToSlice convert Iter to slice
func (si *sliceIter[A]) ToSlice() []A {
	out := make([]A, len(si.slice), len(si.slice))
	_ = copy(out, si.slice)
	return out
}

// Take : take n elements of SliceIter
// if n <= size => take n elements
// if n > size => take up to the end of the SliceIter
func (si *sliceIter[A]) Take(n int) SliceIter[A] {
	var outSlice []A
	for si.HasNext() && si.current-1 <= n {
		outSlice = append(outSlice, si.Next())
	}
	return &sliceIter[A]{
		slice:   outSlice,
		size:    len(outSlice),
		current: 0,
	}
}

// Filter filters SliceIter based on predicate and return new SliceIter
func (si *sliceIter[A]) Filter(fn func(value A) bool) SliceIter[A] {
	var out []A
	for si.HasNext() {
		if value := si.Next(); fn(value) {
			out = append(out, value)
		}
	}
	return &sliceIter[A]{
		slice:   out,
		size:    len(out),
		current: 0,
	}
}

// Map maps F: A => B any: sliceIter[any]
func (si *sliceIter[A]) Map(fn func(value A) any) Iter[any] {
	res := Map[A, any](si, fn)
	return res
}

// Reduce consume the iterator and apply the reduce function
func (si *sliceIter[A]) Reduce(fn func(A, A) A) A {
	var result A
	for si.HasNext() {
		value := si.Next()
		result = fn(result, value)
	}
	return result
}

// Fold consume the iterator and apply the fold function
// it behaves like reduce
func (si *sliceIter[A]) Fold(fn func(A, A) A) A {
	return si.Reduce(fn)
}

// FoldLeft consume the iterator and apply the fold function
// it behaves like reduce
func (si *sliceIter[A]) FoldLeft(initialValue any, fn func(any, A) any) any {
	for si.HasNext() {
		value := si.Next()
		initialValue = fn(initialValue, value)
	}
	return initialValue
}

// Foreach F: A => for all element of the Iter apply side affect function
func (si *sliceIter[A]) Foreach(fn func(A)) {
	for si.HasNext() {
		fn(si.Next())
	}
}

// Slice Creates an iterator returning an interval of the values produced by this iterator.
func (si *sliceIter[A]) Slice(from, until int) SliceIter[A] {
	// from is beyond the end of the Iter or from is negative
	if from > si.size || from < 0 {
		return &sliceIter[A]{
			slice:   make([]A, 0),
			size:    0,
			current: 0,
		}
	}
	// from > until or until is negative
	if from > until {
		return &sliceIter[A]{
			slice:   make([]A, 0),
			size:    0,
			current: 0,
		}
	}
	// happy path
	originalSlice := si.slice
	var tempSlice []A
	index := from
	if until <= si.size {
		for ; index <= until; index++ {
			tempSlice = append(tempSlice, originalSlice[index])
		}
		return &sliceIter[A]{
			slice:   tempSlice,
			size:    len(tempSlice),
			current: 0,
		}
	}

	for ; index < si.size; index++ {
		tempSlice = append(tempSlice, originalSlice[index])
	}
	return &sliceIter[A]{
		slice:   tempSlice,
		size:    len(tempSlice),
		current: 0,
	}
}

// Clone copy SliceIter to another SliceIter
func (si *sliceIter[A]) Clone() SliceIter[A] {
	slice := make([]A, si.size)
	copy(slice, si.slice)
	return &sliceIter[A]{
		slice:   slice,
		current: si.current,
		size:    si.size,
	}
}

// Drop :drop n elements of the SliceIter and new SliceIter
func (si *sliceIter[A]) Drop(n int) SliceIter[A] {

	if n < 0 || n >= si.size {
		return &sliceIter[A]{
			slice:   make([]A, 0, 0),
			size:    0,
			current: 0,
		}
	}

	var slice []A
	var index int
	for si.HasNext() {
		if index >= n {
			slice = append(slice, si.Next())
			index++
			continue
		}
		index++
		si.Next()
	}
	return &sliceIter[A]{
		slice:   slice,
		size:    len(slice),
		current: 0,
	}
}

// Contains return True if element exists
func (si *sliceIter[A]) Contains(elm A) bool {
	for si.HasNext() {
		value := si.Next()
		if any(value) == any(elm) {
			return true
		}
	}
	return false
}

func (si *sliceIter[A]) ToIter() Iter[A] {
	return any(si).(Iter[A])
}
