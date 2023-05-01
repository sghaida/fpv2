// Package collections ...
package collections

// List is generic slice
type List[A any] []A

// Reducer defines the shape of the reduce function
type Reducer[A, B any] func(acc B, value A) B

// Mapper defines the shape of the Map function
type Mapper[A, B any] func(elm A) B

// ListOps provide the slice| array operations
type ListOps[A any] interface {
	Size() int
	Foreach(fn func(elm A))
	Append(elm A) List[A]
	Prepend(elm A) List[A]
	Take(n int) List[A]
}

// Size return size
func (a List[A]) Size() int {
	return len(a)
}

// Foreach apply fn for each element in A
func (a List[A]) Foreach(fn func(elm A)) {
	for _, value := range a {
		fn(value)
	}
}

// Append appends element to the List
func (a List[A]) Append(elm A) List[A] {
	arr := make(List[A], len(a), len(a)+1)
	copy(arr, a)
	arr = append(arr, elm)
	return arr
}

// Prepend prepends element to the List
func (a List[A]) Prepend(elm A) List[A] {
	arr := make(List[A], 0, len(a)+1)
	arr = append(arr, elm)
	arr = append(arr, a...)
	return arr
}

// Take takes n element from the slice
func (a List[A]) Take(n int) List[A] {
	if n < 0 {
		return make(List[A], 0)
	}
	arr := make(List[A], 0, n)
	for i := 0; i < n && i < a.Size(); i++ {
		arr = append(arr, a[i])
	}
	return arr
}

// Map apply F: A => B: List[B]
func Map[A, B any](lst List[A], mapper Mapper[A, B]) List[B] {
	acc := make([]B, 0, len(lst))
	for _, elm := range lst {
		acc = append(acc, mapper(elm))
	}
	return acc
}

// Reduce Apply F: A => B: B
func Reduce[A, B any](lst List[A], reducer Reducer[A, B]) B {
	var acc B
	return FoldLeft(lst, acc, reducer)
}

// FoldLeft applies F: (B, A) => B: B
func FoldLeft[A, B any](lst List[A], acc B, fn Reducer[A, B]) B {
	for _, value := range lst {
		acc = fn(acc, value)
	}
	return acc
}
