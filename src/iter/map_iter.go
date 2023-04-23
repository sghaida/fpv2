// Package iter ...
package iter

// MapIter list the operations that is supported by the MapIter
type MapIter[A comparable, B any] interface {
	Count() int
	Size() int
	HasNext() bool
	Next() MapEntry[A, B]
	ToMap() map[A]B
}

// MapEntry holds the values of map Key Value
type MapEntry[K comparable, V any] struct {
	Key K
	Val V
}

type mapIter[A comparable, B any] struct {
	m    map[A]B
	iter SliceIter[MapEntry[A, B]]
}

// FromMap Converts the Map to Iter
func FromMap[A comparable, B any](from map[A]B) MapIter[A, B] {
	var entries []MapEntry[A, B]
	for k, v := range from {
		entries = append(entries, MapEntry[A, B]{Key: k, Val: v})
	}
	return &mapIter[A, B]{
		m:    from,
		iter: FromSlice(entries),
	}
}

// HasNext check if there is next element
func (mi *mapIter[A, B]) HasNext() bool {
	return mi.iter.HasNext()
}

// Next return the next element in the slice if available
func (mi *mapIter[A, B]) Next() MapEntry[A, B] {
	return mi.iter.Next()

}

// Count return the size of the iter and move to the end of the iter
func (mi *mapIter[A, B]) Count() int {
	return mi.iter.Count()
}

// Size return the size of the iter
func (mi *mapIter[A, B]) Size() int {
	return mi.iter.Size()
}

// ToMap builds a map from an Iter.
func (mi *mapIter[A, B]) ToMap() map[A]B {
	out := map[A]B{}
	for mi.iter.HasNext() {
		value := mi.iter.Next()
		out[value.Key] = value.Val
	}
	return out
}
