package iter

// EmptyIter definition of EmptyIter
type EmptyIter[A any] interface {
	Iter[A]
	Size() int
}

type emptyIter[A any] struct{}

// Empty creates empty Iter
func Empty[A any]() EmptyIter[A] {
	return emptyIter[A]{}
}

// HasNext check if there is next element
func (e emptyIter[A]) HasNext() bool {
	return false
}

// Next return the zero value of the type with false, since its empty
func (e emptyIter[A]) Next() A {
	var defaultValue A
	return defaultValue
}

// Count return the count which is zero
func (e emptyIter[A]) Count() int {
	return 0
}

// Size return the size which is zero
func (e emptyIter[A]) Size() int {
	return 0
}
