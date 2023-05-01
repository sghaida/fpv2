// Package iter ...
package iter

// MapOps include the operations that can be done on a MapIter
type MapOps[A comparable, B any] interface {
	Clone() MapIter[A, B]
	Contains(elm A) bool
	Filter(fn func(A) bool) MapIter[A, B]
	Fold(fn func(acc B, key A, value B) B) B
	FoldLeft(acc any, fn func(key A, value B) any) any
	Foreach(fn func(key A, value B))
	Map(fn func(key A, value B) any) MapIter[A, any]
	Reduce(fn func(v1, v2 B) B) B
	ToSlice() []MapEntry[A, B]
}

// MapIter list the operations that is supported by the MapIter
type MapIter[A comparable, B any] interface {
	Count() int
	Size() int
	HasNext() bool
	Next() MapEntry[A, B]
	ToMap() map[A]B
	MapOps[A, B]
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
	value := mi.iter.Next()
	delete(mi.m, value.Key)
	return value
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

// Clone clones MapIter
func (mi *mapIter[A, B]) Clone() MapIter[A, B] {
	m := make(map[A]B)
	for k, v := range mi.m {
		m[k] = v
	}
	return FromMap(m)
}

// Contains check if key exists
func (mi *mapIter[A, B]) Contains(key A) bool {
	_, ok := mi.m[key]
	return ok
}

// Filter filters a map based on key predicate
func (mi *mapIter[A, B]) Filter(fn func(A) bool) MapIter[A, B] {
	m := make(map[A]B)
	for k, v := range mi.m {
		if fn(k) {
			m[k] = v
		}
	}
	return FromMap(m)
}

// Reduce consume the iterator and apply the reduce function
func (mi *mapIter[A, B]) Reduce(fn func(v1, v2 B) B) B {
	var result B
	for mi.HasNext() {
		entry := mi.Next()
		result = fn(result, entry.Val)
	}
	return result
}

// Fold consume the iterator and apply the fold function
func (mi *mapIter[A, B]) Fold(fn func(acc B, key A, value B) B) B {
	var acc B
	for mi.HasNext() {
		value := mi.Next()
		acc = fn(acc, value.Key, value.Val)
	}
	return acc
}

// FoldLeft consume the iterator and apply the fold function
func (mi *mapIter[A, B]) FoldLeft(acc any, fn func(key A, value B) any) any {
	for mi.HasNext() {
		value := mi.Next()
		acc = fn(value.Key, value.Val)
	}
	return acc
}

// Foreach F: A => for all element of the Iter apply side affect function
func (mi *mapIter[A, B]) Foreach(fn func(key A, value B)) {
	for mi.HasNext() {
		value := mi.Next()
		fn(value.Key, value.Val)
	}
}

// Map maps F: A, B => any
func (mi *mapIter[A, B]) Map(fn func(key A, value B) any) MapIter[A, any] {
	m := make(map[A]any)
	for mi.HasNext() {
		value := mi.Next()
		entry := fn(value.Key, value.Val)
		m[value.Key] = entry
	}
	return FromMap(m)
}

// ToSlice Convert MapIter to Slice
func (mi *mapIter[A, B]) ToSlice() []MapEntry[A, B] {
	return mi.iter.ToSlice()
}
