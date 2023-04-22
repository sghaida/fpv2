package iter

// SliceOps include the operations that can be done on a SliceIter
type SliceOps[A any] interface {
	ToSlice() []A
	Take(n int) SliceIter[A]
	Filter(fn func(A) bool) SliceIter[A]
	Map(fn func(A) any) Iter[any]
	Reduce(fn func(A, A) A) A
	Fold(fn func(A, A) A) A
	FoldLeft(initialValue any, fn func(any, A) any) any
	Foreach(fn func(A))
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

// Take : take n elements of SliceIter
// if n <= size => take n elements
// if n > size => take up to the end of the SliceIter
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

// Filter filters SliceIter based on predicate and return new SliceIter
func (iter *sliceIter[A]) Filter(fn func(value A) bool) SliceIter[A] {
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
func (iter *sliceIter[A]) Map(fn func(value A) any) Iter[any] {
	res := Map[A, any](iter, fn)
	return res
}

// Reduce consume the iterator and apply the reduce function
func (iter *sliceIter[A]) Reduce(fn func(A, A) A) A {
	var result A
	for iter.HasNext() {
		value := iter.Next()
		result = fn(result, value)
	}
	return result
}

// Fold consume the iterator and apply the fold function
// it behaves like reduce
func (iter *sliceIter[A]) Fold(fn func(A, A) A) A {
	return iter.Reduce(fn)
}

// FoldLeft consume the iterator and apply the fold function
// it behaves like reduce
func (iter *sliceIter[A]) FoldLeft(initialValue any, fn func(any, A) any) any {
	for iter.HasNext() {
		value := iter.Next()
		initialValue = fn(initialValue, value)
	}
	return initialValue
}

// Foreach F: A => for all element of the Iter apply side affect function
func (iter *sliceIter[A]) Foreach(fn func(A)) {
	for iter.HasNext() {
		fn(iter.Next())
	}
}
