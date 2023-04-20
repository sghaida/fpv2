package iter

import "errors"

// RangeIter definition of RangeIter
type RangeIter[A Number] interface {
	Size() int
	Iter[A]
}

type rangeIter[A Number] struct {
	start, end, step, size A
}

// Range creates Range Iter
// on success => return the Iter
// on failure => return the error
func Range[A Number](start, end, step A) (RangeIter[A], error) {
	if end < start {
		return emptyIter[A]{}, errors.New("end < start")
	}
	if step < 0 {
		return emptyIter[A]{}, errors.New("step < 0")
	}
	return &rangeIter[A]{
		start: start,
		end:   end,
		step:  step,
		size:  end - start/step,
	}, nil
}

// HasNext check if there is next element
func (iter *rangeIter[A]) HasNext() bool {
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
	return loc
}

// Count return the size of the iter
func (iter *rangeIter[A]) Count() int {
	count := (iter.end - iter.start) / iter.step
	iter.start = iter.end
	return int(count)
}

// Size return the size of the iter
func (iter *rangeIter[A]) Size() int {
	return int(iter.size)
}
