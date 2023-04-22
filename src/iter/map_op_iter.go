// Package iter ...
package iter

// MapIter interface wraps basic Iter
type MapIter[A, B any] interface {
	Iter[A]
}

type mapIter[A, B any] struct {
	from  Iter[A]
	mapFn func(A) B
}

// Map create a MapIter
func Map[A, B any](iter Iter[A], fn func(A) B) Iter[B] {
	return &mapIter[A, B]{
		from:  iter,
		mapFn: fn,
	}
}

// HasNext check if there is next element
func (iter *mapIter[A, B]) HasNext() bool {
	if iter.from.HasNext() {
		return true
	}
	return false
}

// Next return the next element in the slice if available
// please note that the default value of type B could be nil
func (iter *mapIter[A, B]) Next() B {
	if !iter.from.HasNext() {
		var zero B
		return zero
	}
	value := iter.from.Next()
	return iter.mapFn(value)
}

// Count return the size of the iter and move to the end of the iter
func (iter *mapIter[A, B]) Count() int {
	return iter.from.Count()
}

// Size return the size of the iter
func (iter *mapIter[A, B]) Size() int {
	return iter.from.Size()
}
