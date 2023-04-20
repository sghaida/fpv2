package iter

// Number allowed numbers as type for the iter
type Number interface {
	~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Iter interface holds 2 methods
// Next => to return current value
// Count => to return the size of the iter
type Iter[A any] interface {
	Next() (A, bool)
	Count() int
}
