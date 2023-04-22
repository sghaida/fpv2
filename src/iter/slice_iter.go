package iter

type SliceIter[A any] interface {
	Iter[A]
	ToSlice() []A
	Take(n int) SliceIter[A]
}

type sliceIter[T any] struct {
	slice   []T
	size    int
	current int
}

// FromSlice creates Iter from slice
func FromSlice[A any](slice []A) SliceIter[A] {
	return &sliceIter[A]{slice: slice, size: len(slice), current: 0}
}

// HasNext check if there is next element
func (iter *sliceIter[A]) HasNext() bool {
	if len(iter.slice) != 0 {
		return true
	}
	return false
}

// Next return the next element in the slice if available
func (iter *sliceIter[A]) Next() A {
	if !iter.HasNext() {
		var zero A
		return zero
	}
	item := iter.slice[0]
	iter.slice = iter.slice[1:]
	iter.current++
	iter.size = len(iter.slice)
	return item
}

// Count return the size of the iter and move to the end of the iter
func (iter *sliceIter[A]) Count() int {
	count := len(iter.slice)
	iter.slice = []A{}
	iter.current = iter.size
	return count
}

// Size return the size of the iter
func (iter *sliceIter[A]) Size() int {
	return iter.size
}

// ToSlice convert Iter to slice
func (iter *sliceIter[A]) ToSlice() []A {
	out := make([]A, len(iter.slice), len(iter.slice))
	_ = copy(out, iter.slice)
	return out
}

func (iter *sliceIter[A]) Take(n int) SliceIter[A] {
	var outSlice []A
	for iter.HasNext() && iter.current-1 <= n {
		outSlice = append(outSlice, iter.Next())
	}
	return &sliceIter[A]{
		slice:   outSlice,
		size:    len(outSlice),
		current: 0,
	}

}
