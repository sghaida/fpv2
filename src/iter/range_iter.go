package iter

import (
	"errors"
)

// RangeIter definition of RangeIter
type RangeIter[A Number] interface {
	Iter[A]
	Size() int
	Take(n A, step A) RangeIter[A]
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
		size:  (end - start) / step,
	}, nil
}

// HasNext check if there is next element
func (iter *rangeIter[A]) HasNext() bool {
	if iter.size <= 0 {
		return false
	}
	return iter.start+iter.step <= iter.end
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
	count := (iter.end - iter.start) / iter.step
	iter.start = iter.end
	iter.size = 0
	return int(count)
}

// Size return the size of the iter
func (iter *rangeIter[A]) Size() int {
	return int(iter.size)
}

// Take :take the first n elements from an Iter with a configured step
// and if the end of the iter > n then take up to the end of the iter
func (iter *rangeIter[A]) Take(n A, step A) RangeIter[A] {
	if iter.start+n <= iter.end {
		return &rangeIter[A]{
			start: iter.start,
			end:   n,
			step:  step,
			size:  (n - iter.start) / step,
		}
	}

	return &rangeIter[A]{
		start: iter.start,
		end:   iter.end,
		step:  step,
		size:  iter.size,
	}
}
