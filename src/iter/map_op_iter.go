// Package iter ...
package iter

// MapOpIter interface wraps basic Iter
type MapOpIter[A, B any] interface {
	Iter[A]
	ToIter() Iter[B]
}

type mapOpIter[A, B any] struct {
	from  Iter[A]
	mapFn func(A) B
}

// Map create a MapOpIter
func Map[A, B any](iter Iter[A], fn func(A) B) Iter[B] {
	return &mapOpIter[A, B]{
		from:  iter,
		mapFn: fn,
	}
}

// HasNext check if there is next element
func (moi *mapOpIter[A, B]) HasNext() bool {
	if moi.from.HasNext() {
		return true
	}
	return false
}

// Next return the next element in the slice if available
// please note that the default value of type B could be nil
func (moi *mapOpIter[A, B]) Next() B {
	if !moi.from.HasNext() {
		var zero B
		return zero
	}
	value := moi.from.Next()
	return moi.mapFn(value)
}

// Count return the size of the iter and move to the end of the iter
func (moi *mapOpIter[A, B]) Count() int {
	return moi.from.Count()
}

// Size return the size of the iter
func (moi *mapOpIter[A, B]) Size() int {
	return moi.from.Size()
}
