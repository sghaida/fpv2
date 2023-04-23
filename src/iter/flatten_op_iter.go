// Package iter ...
package iter

// Flatten applies a function to all items of the specified iterator, recurse and concatenate The resulting iterators
func Flatten[A any](iter Iter[Iter[A]]) Iter[A] {
	return &flattenIterator[A]{iter: iter}
}

type flattenIterator[A any] struct {
	iter Iter[Iter[A]]
	head Iter[A]
}

func (fi *flattenIterator[A]) Next() A {
	for {
		if fi.head == nil {
			if !fi.HasNext() {
				var zero A
				return zero
			}
			fi.head = fi.iter.Next()
		}
		if fi.HasNext() {
			return fi.Next()
		}
		fi.head = nil
	}
}

func (fi *flattenIterator[A]) HasNext() bool {
	if fi.head != nil {
		return fi.HasNext()
	}
	return fi.iter.HasNext()
}

func (fi *flattenIterator[A]) Count() int {
	if fi.head != nil {
		return fi.head.Count()
	}
	return fi.Count()
}

func (fi *flattenIterator[A]) Size() int {
	if fi.head != nil {
		return fi.head.Size()
	}
	return fi.iter.Size()
}
