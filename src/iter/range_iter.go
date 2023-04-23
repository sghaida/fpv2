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
func (iter *rangeIter[A]) HasNext() bool {
	if iter.size <= 0 {
		return false
	}
	return iter.start+iter.step <= iter.end+1
}

// Next return the current step in the Iter
// on success => current step in the Iter, true
// on failure => return zero value of the Type and false
func (iter *rangeIter[A]) Next() A {
	if !iter.HasNext() {
		var value A
		return value
	}
	loc := iter.start
	iter.start += iter.step
	iter.size--
	return loc
}

// Count return the size of the iter
func (iter *rangeIter[A]) Count() int {
	if iter.size <= 0 {
		return 0
	}
	count := ((iter.end - iter.start) / iter.step) + 1
	iter.start = iter.end
	iter.size = 0
	return int(count)
}

// Size return the size of the iter
func (iter *rangeIter[A]) Size() int {
	return int(iter.size)
}

// Take :take the first n elements iter an Iter with a configured step
// and if the end of the iter > n then take up to the end of the iter
func (iter *rangeIter[A]) Take(n A, step A) RangeIter[A] {
	if iter.start+n <= iter.end {
		return &rangeIter[A]{
			start: iter.start,
			end:   n,
			step:  step,
			size:  n / step,
		}
	}

	return &rangeIter[A]{
		start: iter.start,
		end:   iter.end,
		step:  step,
		size:  iter.size,
	}
}

// Filter filters SliceIter based on predicate and return new SliceIter
func (iter *rangeIter[A]) Filter(fn func(value A) bool) SliceIter[A] {
	var out []A
	for iter.HasNext() {
		if value := iter.Next(); fn(value) {
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
func (iter *rangeIter[A]) Map(fn func(value A) any) Iter[any] {
	res := Map[A, any](iter, fn)
	return res
}

// Reduce consume the iterator and apply the reduce function
func (iter *rangeIter[A]) Reduce(fn func(A, A) A) A {
	var result A
	for iter.HasNext() {
		value := iter.Next()
		result = fn(result, value)
	}
	return result
}

// Fold consume the iterator and apply the fold function
// it behaves like reduce
func (iter *rangeIter[A]) Fold(fn func(A, A) A) A {
	return iter.Reduce(fn)
}

// FoldLeft consume the iterator and apply the fold function
// it behaves like reduce
func (iter *rangeIter[A]) FoldLeft(initialValue any, fn func(any, A) any) any {
	for iter.HasNext() {
		value := iter.Next()
		initialValue = fn(initialValue, value)
	}
	return initialValue
}

// Foreach F: A => for all element of the Iter apply side affect function
func (iter *rangeIter[A]) Foreach(fn func(A)) {
	for iter.HasNext() {
		fn(iter.Next())
	}
}

// Slice Creates an iterator returning an interval of the values produced by this iterator.
func (iter *rangeIter[A]) Slice(from, until A) SliceIter[A] {
	// iter is beyond the end of the Iter or iter is negative
	if from < 0 || until < 0 {
		return &sliceIter[A]{
			slice:   make([]A, 0),
			size:    0,
			current: 0,
		}
	}

	// happy path
	originalIter := iter
	var tempSlice []A

	for originalIter.HasNext() {
		if iter.start < from {
			iter.Next()
			continue
		}
		if iter.start > until || iter.start == iter.end {
			return &sliceIter[A]{
				slice:   tempSlice,
				size:    len(tempSlice),
				current: 0,
			}
		}
		value := iter.Next()
		tempSlice = append(tempSlice, value)
	}

	return &sliceIter[A]{
		slice:   make([]A, 0),
		size:    0,
		current: 0,
	}
}

// Clone copy SliceIter to another SliceIter
func (iter *rangeIter[A]) Clone() RangeIter[A] {
	return &rangeIter[A]{
		start: iter.start,
		end:   iter.end,
		step:  iter.step,
		size:  iter.size,
	}
}

// Drop :drop n elements of the SliceIter and new SliceIter
func (iter *rangeIter[A]) Drop(n A) RangeIter[A] {

	for iter.HasNext() {
		if iter.start > n {
			return iter
		}
		iter.Next()
	}
	return &rangeIter[A]{
		start: 0,
		end:   0,
		step:  iter.step,
		size:  0,
	}
}

// Contains return True if element exists
func (iter *rangeIter[A]) Contains(elm A) bool {
	for iter.HasNext() {
		value := iter.Next()
		if any(value) == any(elm) {
			return true
		}
	}
	return false
}

func (iter *rangeIter[A]) ToIter() Iter[A] {
	return any(iter).(Iter[A])
}

// ToSlice convert Iter to slice
func (iter *rangeIter[A]) ToSlice() []A {
	size := any(iter.size).(int)
	out := make([]A, 0, size)
	for iter.HasNext() {
		out = append(out, iter.Next())
	}
	return out
}
