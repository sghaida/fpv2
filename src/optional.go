// Package src ...
package src

import (
	"errors"
	"fmt"
	"github.com/sghaida/fpv2/src/utils"
)

// Type : SomeType | NoneType
type Type string

const (
	// SomeType : Some
	SomeType Type = "Some"
	// NoneType : None
	NoneType Type = "None"
)

// ErrorNoneValue : ErrorNoneValue
var ErrorNoneValue = errors.New("NoneValue presented")

// ErrorSomeValue : ErrorSomeValue
var ErrorSomeValue = errors.New("SomeValue presented")

// Option : Option type
type Option[A any] struct {
	value A
	t     Type
}

// NewOptional create new Option
func NewOptional[A any](value A) Option[A] {
	if utils.IsNilOrZeroValue(value) || utils.IsPtr(value) {
		return Option[A]{t: NoneType}
	}
	return Option[A]{value: value, t: SomeType}
}

// Some will create Option of SomeType in case the value is not Nil
// and returns error with NoneValue in case the Value is Nil
func Some[A any](value A) (Option[A], error) {
	if utils.IsNilOrZeroValue(value) || utils.IsPtr(value) {
		return None[A](), ErrorNoneValue
	}
	return Option[A]{value: value, t: SomeType}, nil
}

// None will create an Option of NoneType
func None[A any]() Option[A] {
	return Option[A]{t: NoneType}
}

// GetType return Option Type
func (o Option[A]) GetType() Type {
	return o.t
}

// IsNone validates if the Option of Type None
func (o Option[A]) IsNone() bool {
	return o.t == NoneType
}

// IsSome validates if the Option of Type Some
func (o Option[A]) IsSome() bool {
	return o.t == SomeType
}

// Get return wrapped value of the Option
func (o Option[A]) Get() A {
	return o.value
}

// Apply applies function f and returns another Option after apply
func (o Option[A]) Apply(f func(value A) A) Option[A] {
	// TODO think about the case where the value is reference type such as slice, Map
	//  as the original value will be update and is this behaviour needed
	//  the alternative is to do deep copy for the reference types which requires reflection
	if o.t == SomeType {
		return Option[A]{
			value: f(o.value),
			t:     SomeType,
		}
	}
	return o
}

// Take return the wrapped value in the context of SomeType, and error in the context of NoneType
func (o Option[A]) Take() (value A, err error) {
	value = o.value
	if o.IsNone() {
		err = ErrorNoneValue
		return
	}
	return
}

// TakeOr return the Value in the context of SomeType or value in the context of NoneType
func (o Option[A]) TakeOr(value A) A {
	if o.IsNone() {
		return value
	}
	return o.value
}

// TakeOrElse return the Value in the context of SomeType and applying function f() and return in the context of NoneType
func (o Option[A]) TakeOrElse(fn func() A) A {
	if o.IsNone() {
		return fn()
	}
	return o.value
}

// Or return Original Option in the context of SomeType or value in the context of NoneType
func (o Option[A]) Or(value Option[A]) Option[A] {
	if o.IsNone() {
		return value
	}
	return o
}

// OrElse return Original Option in the context of SomeType or return f()  in the context of NoneType
func (o Option[A]) OrElse(fn func() Option[A]) Option[A] {
	if o.IsNone() {
		return fn()
	}
	return o
}

// IfSome execute f() if the Option is SomeType
func (o Option[A]) IfSome(f func(value A) A) (A, error) {
	// TODO think about the case where the value is reference type such as slice, Map
	//  as the original value will be update and is this behaviour needed
	//  the alternative is to do deep copy for the reference types which requires reflection
	if o.IsSome() {
		return f(o.value), nil
	}
	return o.value, ErrorNoneValue
}

// IfNone execute f() if the Option is NoneType
func (o Option[A]) IfNone(fn func() A) (A, error) {
	if o.IsNone() {
		return fn(), nil
	}
	return o.value, ErrorSomeValue
}

// Unwrap return the underlying value irrelevant if it is SomeType or NoneType
func (o Option[T]) Unwrap() T {
	return o.Get()
}

// Flatten :Flatten an Option[Option[A]] --> Option[A]
func (o Option[A]) Flatten() Option[A] {
	// if the inner type is also an option Flatten it
	if o, ok := any(o.Get()).(Option[A]); ok {
		return o
	}
	return o
}

// String return String representation of the Option
func (o Option[A]) String() string {
	if o.IsNone() {
		return fmt.Sprintf("%s", NoneType)
	}
	if stringer, ok := interface{}(o.value).(fmt.Stringer); ok {
		return fmt.Sprintf("Some[%s]", stringer)
	}
	return fmt.Sprintf("Some[%v]", o.value)
}

// Map for an Option[A] apply Mapper for function A --> Any and return Option[Any]
func (o Option[A]) Map(mapper utils.Mapper[A, any]) Option[any] {
	return Map(o, mapper)
}

// Map for an Option[A] apply Mapper for function A --> B and return Option[B]
func Map[A, B any](option Option[A], mapper utils.Mapper[A, B]) Option[B] {
	if option.IsNone() {
		return None[B]()
	}
	// error handling is skipped as it is being handled by the above check
	someOpt, _ := Some[B](mapper(option.value))

	return someOpt
}

// OptionFlatMapperFn is a function that is applies on type A and return Option[B]
type OptionFlatMapperFn[A, B any] func(A) Option[B]

// FlatMap for Option[A] apply mapper function from A--> Option[B] and return Option[B]
// This function could panic if the mapper is not applicable on A such as in the context of
// Option[Option[Option[A]]
func (o Option[A]) FlatMap(fn OptionFlatMapperFn[A, any]) Option[any] {
	return OptionFlatMap(o, fn)
}

// OptionFlatMap for Option[A] apply mapper function from A--> Option[B] and return Option[B]
// This function could panic if the mapper is not applicable on A such as in the context of
// Option[Option[Option[A]]
func OptionFlatMap[A, B any](option Option[A], mapper OptionFlatMapperFn[A, B]) Option[B] {
	if option.IsNone() {
		return None[B]()
	}
	switch value := any(option.value).(type) {
	case Option[A]:
		v, ok := any(value).(Option[A])
		if ok {
			if v.IsNone() {
				return None[B]()
			}
			return mapper(v.Get())
		}
	case A:
		return mapper(value)
	}
	return None[B]()
}

// ToEither Converts Option to Either if None => Left, if Some => Right
func (o Option[A]) ToEither() Either[any, A] {
	if o.IsNone() {
		return Left[any, A](ErrorLeftValue)
	}
	return Right[any, A](o.value)
}
