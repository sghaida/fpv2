// Package iter ...
package iter

import (
	"errors"
)

type RangeOps[A any] interface {
	Contains(elm A) bool
	Filter(fn func(A) bool) SliceIter[A]
	Fold(fn func(A, A) A) A
	FoldLeft(initialValue any, fn func(any, A) any) any
	Foreach(fn func(A))
	Map(fn func(A) any) Iter[any]
	Reduce(fn func(A, A) A) A
	ToSlice() []A
	ToIter() Iter[A]
	Slice(from, until A) SliceIter[A]
}

type RangeNumberOps[A Number] interface {
	Clone() RangeIter[A]
	Take(n, step A) RangeIter[A]
	Drop(n A) RangeIter[A]
}

// RangeIter definition of RangeIter
type RangeIter[A Number] interface {
	Iter[A]
	RangeOps[A]
	RangeNumberOps[A]
}

type rangeIter[A Number] struct {
	start, end, step, size A
}

// Range creates Range Iter
// on success => return the Iter
// on failure => return the error
func Range[A Number](start, end, step A) (RangeIter[A], error) {
	if end < start {
		return &rangeIter[A]{}, errors.New("end < start")
	}
	if step < 0 {
		return &rangeIter[A]{}, errors.New("step < 0")
	}
	return &rangeIter[A]{
		start: start,
		end:   end,
		step:  step,
		size:  ((end - start) / step) + 1,
	}, nil
}

// HasNext check if there is next element
func (ri *rangeIter[A]) HasNext() bool {
	if ri.size <= 0 {
		return false
	}
	return ri.start+ri.step <= ri.end+1
}

// Next return the current step in the Iter
// on success => current step in the Iter, true
// on failure => return zero value of the Type and false
func (ri *rangeIter[A]) Next() A {
	if !ri.HasNext() {
		var value A
		return value
	}
	loc := ri.start
	ri.start += ri.step
	ri.size--
	return loc
}

// Count return the size of the iter
func (ri *rangeIter[A]) Count() int {
	if ri.size <= 0 {
		return 0
	}
	count := ((ri.end - ri.start) / ri.step) + 1
	ri.start = ri.end
	ri.size = 0
	return int(count)
}

// Size return the size of the iter
func (ri *rangeIter[A]) Size() int {
	return int(ri.size)
}

// Take :take the first n elements iter an Iter with a configured step
// and if the end of the iter > n then take up to the end of the iter
func (ri *rangeIter[A]) Take(n A, step A) RangeIter[A] {
	if ri.start+n <= ri.end {
		return &rangeIter[A]{
			start: ri.start,
			end:   n,
			step:  step,
			size:  n / step,
		}
	}

	return &rangeIter[A]{
		start: ri.start,
		end:   ri.end,
		step:  step,
		size:  ri.size,
	}
}

// Filter filters SliceIter based on predicate and return new SliceIter
func (ri *rangeIter[A]) Filter(fn func(value A) bool) SliceIter[A] {
	var out []A
	for ri.HasNext() {
		if value := ri.Next(); fn(value) {
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
func (ri *rangeIter[A]) Map(fn func(value A) any) Iter[any] {
	res := Map[A, any](ri, fn)
	return res
}

// Reduce consume the iterator and apply the reduce function
func (ri *rangeIter[A]) Reduce(fn func(A, A) A) A {
	var result A
	for ri.HasNext() {
		value := ri.Next()
		result = fn(result, value)
	}
	return result
}

// Fold consume the iterator and apply the fold function
// it behaves like reduce
func (ri *rangeIter[A]) Fold(fn func(A, A) A) A {
	return ri.Reduce(fn)
}

// FoldLeft consume the iterator and apply the fold function
// it behaves like reduce
func (ri *rangeIter[A]) FoldLeft(initialValue any, fn func(any, A) any) any {
	for ri.HasNext() {
		value := ri.Next()
		initialValue = fn(initialValue, value)
	}
	return initialValue
}

// Foreach F: A => for all element of the Iter apply side affect function
func (ri *rangeIter[A]) Foreach(fn func(A)) {
	for ri.HasNext() {
		fn(ri.Next())
	}
}

// Slice Creates an iterator returning an interval of the values produced by this iterator.
func (ri *rangeIter[A]) Slice(from, until A) SliceIter[A] {
	// ri is beyond the end of the Iter or ri is negative
	if from < 0 || until < 0 {
		return &sliceIter[A]{
			slice:   make([]A, 0),
			size:    0,
			current: 0,
		}
	}

	// happy path
	originalIter := ri
	var tempSlice []A

	for originalIter.HasNext() {
		if ri.start < from {
			ri.Next()
			continue
		}
		if ri.start > until || ri.start == ri.end {
			return &sliceIter[A]{
				slice:   tempSlice,
				size:    len(tempSlice),
				current: 0,
			}
		}
		value := ri.Next()
		tempSlice = append(tempSlice, value)
	}

	return &sliceIter[A]{
		slice:   make([]A, 0),
		size:    0,
		current: 0,
	}
}

// Clone copy SliceIter to another SliceIter
func (ri *rangeIter[A]) Clone() RangeIter[A] {
	return &rangeIter[A]{
		start: ri.start,
		end:   ri.end,
		step:  ri.step,
		size:  ri.size,
	}
}

// Drop :drop n elements of the SliceIter and new SliceIter
func (ri *rangeIter[A]) Drop(n A) RangeIter[A] {

	for ri.HasNext() {
		if ri.start > n {
			return ri
		}
		ri.Next()
	}
	return &rangeIter[A]{
		start: 0,
		end:   0,
		step:  ri.step,
		size:  0,
	}
}

// Contains return True if element exists
func (ri *rangeIter[A]) Contains(elm A) bool {
	for ri.HasNext() {
		value := ri.Next()
		if any(value) == any(elm) {
			return true
		}
	}
	return false
}

func (ri *rangeIter[A]) ToIter() Iter[A] {
	return any(ri).(Iter[A])
}

// ToSlice convert Iter to slice
func (ri *rangeIter[A]) ToSlice() []A {
	size := any(ri.size).(int)
	out := make([]A, 0, size)
	for ri.HasNext() {
		out = append(out, ri.Next())
	}
	return out
}
