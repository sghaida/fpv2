package src

import (
	"errors"
	"fp/src/utils"
)

// TODO implement test

// typeSide identifies if Either is left or isRight
type typeSide string

const (
	isLeftSided  typeSide = "left"
	isRightSided typeSide = "right"
)

var (
	// ErrorLeftValue if left value is being passed to right
	ErrorLeftValue  = errors.New("left Value is presented")
	ErrorRightValue = errors.New("right Value is presented")
)

// Either is composite type contains left[A] and right[B]
// left usually used to propagate errors, while right to propagate values
// Please note Either is right biased so operations such as Map and OptionFlatMap will be applied on the right
// and if it is left value then the left will return unchanged
type Either[A, B any] struct {
	left  A
	right B
	side  typeSide
}

// Left create Either from Left value
func Left[A, B any](value A) Either[A, B] {
	return Either[A, B]{
		left: value,
		side: isLeftSided,
	}
}

// Right create Either from Right value if value is presented and Left if Nil or a pointer
func Right[A, B any](value B) Either[A, B] {

	if utils.IsNilOrZeroValue(value) || utils.IsPtr(value) {
		return Either[A, B]{
			left: any(ErrorLeftValue).(A),
			side: isLeftSided,
		}
	}
	return Either[A, B]{
		right: value,
		side:  isRightSided,
	}
}

// IsLeft return True if value is Left
func (e Either[A, B]) IsLeft() bool {
	if e.side == isLeftSided {
		return true
	}
	return false
}

// IsRight return True if value is Right
func (e Either[A, B]) IsRight() bool {
	if e.side == isRightSided {
		return true
	}
	return false
}

// Unwrap return the inner values of Right and Left
// the reason behind that is following the Go idioms since left (A) usually contains errors
// value, err := e.Unwrap()
func (e Either[A, B]) Unwrap() (B, A) {
	return e.right, e.left
}

// Swap swaps right and left values
func (e Either[A, B]) Swap() Either[B, A] {
	side := isRightSided
	if e.side == isRightSided {
		side = isLeftSided
	}

	return Either[B, A]{
		left:  e.right,
		right: e.left,
		side:  side,
	}
}

// TakeLeft Return Left if presented and error if not
func (e Either[A, B]) TakeLeft() (A, error) {
	if e.IsLeft() {
		return e.left, nil
	}

	var v A
	return v, ErrorRightValue
}

// TakeRight Return Right if presented and error if not
func (e Either[A, B]) TakeRight() (B, error) {
	if e.IsRight() {
		return e.right, nil
	}
	var v B
	return v, ErrorLeftValue
}

// Take Return Right if presented and error if not
func (e Either[A, B]) Take() (B, error) {
	return e.TakeRight()
}

// TakeOr Return Right if presented or value if not
func (e Either[A, B]) TakeOr(value B) B {
	if e.IsRight() {
		return e.right
	}
	return value
}

// TakeOrElse Return Right if presented or the result of fn() if not
func (e Either[A, B]) TakeOrElse(fn func() B) B {
	if e.IsRight() {
		return e.right
	}
	return fn()
}

// Or Return Right if presented or either if not
func (e Either[A, B]) Or(either Either[A, B]) Either[A, B] {
	if e.IsRight() {
		return e
	}
	return either
}

// OrElse Return Right if presented or results of fn() if not
func (e Either[A, B]) OrElse(fn func() Either[A, B]) Either[A, B] {
	if e.IsRight() {
		return e
	}
	return fn()
}

// Map : map is right biased if the value is left it will return unchanged
// and if the value is right mapper will be applied
// and new either will return contain the right changed after applying the mapper
func (e Either[A, B]) Map(mapper utils.Mapper[B, any]) Either[A, any] {
	switch e.side {
	case isLeftSided:
		return Left[A, any](e.left)
	default:
		value := mapper(e.right)
		return Right[A, any](value)
	}
}

// Fold takes to fn f1: A => C , f2: B => C and applies f1 in case off Left and f2 in case of Right
// and returns the resulting value
func (e Either[A, B]) Fold(a2c utils.Mapper[A, any], b2c utils.Mapper[B, any]) any {
	return foldEither(e, a2c, b2c)
}

// foldEither takes to fn f1: A => C , f2: B => C and applies f1 in case off Left and f2 in case of Right
// and returns the resulting value
func foldEither[A, B, C any](either Either[A, B], a2c utils.Mapper[A, C], b2c utils.Mapper[B, C]) C {
	switch either.side {
	case isLeftSided:
		return a2c(either.left)
	default:
		return b2c(either.right)
	}
}

// Exists validate if value exists on Right using some Predicate
func (e Either[A, B]) Exists(fn utils.Predicate[B]) bool {
	if e.side == isLeftSided {
		return false
	}
	return fn(e.right)
}

type flatten[A, B any] func(value B) Either[A, any]

// FlatMapperFn function definition that takes Right value and apply the function
type FlatMapperFn[A, B, C any] func(value B) Either[A, C]

// EitherFlatMapFn function definition that takes Right value and apply the function
type EitherFlatMapFn[A, B, C any] func(value B) Either[A, C]

// FlatMap takes an Either and function that applies on Right
// if the Either IsRight then it returns Either of Right and if Left returns the Either of Left as Is
// if OptionFlatMap is not applicable will return Left
// e.g.
//
//	Either[A, B] right => Either[any, B] right
//	Either[A, Either[B, C]] right -> right  => Either[any, C] right
//	Either[A, Either[B, C]] right -> left  => Either[B, any] left
//	Either[A, B] left => Either [A, C] left
//
// this function will panic if you get more than 2 levels of Either
// for such a case please use Flatten multiple times
func (e Either[A, B]) FlatMap(fn EitherFlatMapFn[A, B, any]) Either[A, any] {
	return EitherFlatMap(e, fn)
}

// EitherFlatMap takes an Either and function that applies on Right
// if the Either IsRight then it returns Either of Right and if Left returns the Either of Left as Is
// if OptionFlatMap is not applicable will return Left
// e.g.
//
//	Either[A, B] right => Either[any, B] right
//	Either[A, Either[B, C]] right -> right  => Either[any, C] right
//	Either[A, Either[B, C]] right -> left  => Either[B, any] left
//	Either[A, B] left => Either [A, C] left
//
// this function will panic if you get more than 2 levels of Either
// for such a case please use Flatten multiple times
func EitherFlatMap[A, B, C any](e Either[A, B], fn EitherFlatMapFn[A, B, C]) Either[A, C] {
	// Evaluate left side
	if e.IsLeft() {
		return Left[A, C](e.left)
	}
	return fn(any(e.right).(B))
}

// ToOption converts to Option if IsLeft => None[any]() , if IsRight then Some[B](value)
func (e Either[A, B]) ToOption() Option[B] {
	if e.IsRight() {
		some, err := Some[B](e.right)
		// can fail due to passing nil or pointers
		if err != nil {
			return None[B]()
		}
		return any(some).(Option[B])
	}
	return None[B]()
}
