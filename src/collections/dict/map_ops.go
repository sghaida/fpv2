package dict

import (
	"github.com/sghaida/fpv2/src/collections/list"
)

// Dict is a type alias for map
type Dict[A comparable, B any] map[A]B

// Ops is an interface that defines a set of operations that can be performed on a Dict of elements of type A.
// The `A` type parameter specifies the type of elements that the Dict can hold.
type Ops[A comparable, B any] interface {
	// Size returns the number of elements in the Dict.
	Size() int
	// Foreach applies the given function `fn` to each element in the Dict.
	Foreach(fn func(key A))
	// Add returns a new Dict that is the result of appending the given `key` to the Dict.
	Add(key A, value B) Dict[A, B]
	// Remove returns a new Dict that is the result of prepending the given element `elm` to the beginning of the current Dict.
	Remove(key A) Dict[A, B]
	// Contains check if Key exists
	Contains(key A) bool
	// Keys return List of keys
	Keys() list.List[A]
	// Values return List of values
	Values() list.List[B]
	// KeysValues extract keys and values
	KeysValues() (list.List[A], list.List[B])
	// Concat results in concatenation of Both Dict and other Dict
	Concat(other Dict[A, B]) Dict[A, B]
	// Clone copy the Dict to another one
	Clone() Dict[A, B]
}

// NewDict create new Dict
func NewDict[A comparable, B any]() Dict[A, B] {
	m := make(map[A]B)
	return m
}

// Size return the Dict size
func (dict Dict[A, B]) Size() int {
	return len(dict)
}

// Contains check if Dict contains key
func (dict Dict[A, B]) Contains(key A) bool {
	if _, ok := dict[key]; ok {
		return true
	}
	return false
}

// Add adds an element to Dict
func (dict Dict[A, B]) Add(key A, value B) Dict[A, B] {
	dict[key] = value
	return dict
}

// Remove remove element from Dict
func (dict Dict[A, B]) Remove(key A) Dict[A, B] {
	if dict.Contains(key) {
		delete(dict, key)
	}
	return dict
}

// Keys extract Dict keys
func (dict Dict[A, B]) Keys() list.List[A] {
	acc := list.List[A]{}
	return FoldLeft(dict, acc, func(acc list.List[A], key A, _ B) list.List[A] {
		return acc.Append(key)
	})
}

// Values extract Dict values
func (dict Dict[A, B]) Values() list.List[B] {
	acc := list.List[B]{}
	return FoldLeft(dict, acc, func(acc list.List[B], _ A, value B) list.List[B] {
		return acc.Append(value)
	})
}

// KeysValues returns a tuple of two lists containing all keys and values in the dictionary.
// The order of the keys and values in the returned lists is undefined.
// The time complexity is O(n), where n is the size of the dictionary.
// The space complexity is O(n), where n is the size of the dictionary.
//
// Example usage:
//
//	dict := NewDict[string, int]()
//	dict.Add("a", 1)
//	dict.Add("b", 2)
//	dict.Add("c", 3)
//	keys, values := dict.KeysValues()
//	fmt.Println(keys)   // [a b c]
//	fmt.Println(values) // [1 2 3]
func (dict Dict[A, B]) KeysValues() (list.List[A], list.List[B]) {
	keys := list.List[A]{}
	values := list.List[B]{}

	acc := struct {
		k list.List[A]
		v list.List[B]
	}{
		k: keys,
		v: values,
	}

	acc = FoldLeft(dict, acc, func(acc struct {
		k list.List[A]
		v list.List[B]
	}, key A, value B) struct {
		k list.List[A]
		v list.List[B]
	} {
		acc.k = acc.k.Append(key)
		acc.v = acc.v.Append(value)
		return acc
	})

	return acc.k, acc.v
}

// Foreach applies fn on each key|value in Dict
func (dict Dict[A, B]) Foreach(fn func(key A, value B)) {
	for k, v := range dict {
		fn(k, v)
	}
}

// Clone copies the Dict to another Dict
func (dict Dict[A, B]) Clone() Dict[A, B] {
	cloned := Dict[A, B]{}
	for key, value := range dict {
		cloned[key] = value
	}
	return cloned
}

// Concat concatenate 2 Dicts by copying both to new Dict
// if key was duplicated on both dicts then one of them will be overridden in the next write
func (dict Dict[A, B]) Concat(other Dict[A, B]) Dict[A, B] {
	if other == nil || other.Size() == 0 {
		return dict
	}
	if dict == nil || dict.Size() == 0 {
		return other
	}

	merged := Dict[A, B]{}
	// copy the original Dict
	for key, value := range dict {
		merged[key] = value
	}
	for key, value := range other {
		merged[key] = value
	}
	return merged
}

// Map applies a function fn to each key-value pair in the input dictionary dict
// and returns a new dictionary with the transformed values.
//
// Time complexity: O(n), where n is the number of key-value pairs in the input dictionary dict.
// Map applies fn to each key-value pair in dict exactly once.
//
// Space complexity: O(n), where n is the number of key-value pairs in the input dictionary dict.
// Map creates a new dictionary with the same number of keys as dict, but with the transformed values.
//
// Example usage:
//
// dict := NewDict[string, int]()
// dict.add("a", 1).dict.add("b", 2).dict.add("c", 3)
// fn := func(_ string, value int) int { return value * 2 }
// result := Map(dict, fn)
func Map[A comparable, B, C any](dict Dict[A, B], fn func(key A, value B) C) Dict[A, C] {
	acc := NewDict[A, C]()
	for key, value := range dict {
		acc[key] = fn(key, value)
	}
	return acc
}

// FoldLeft applies the fn function to an accumulator and each element of the Dict, starting from the leftmost key-value pair.
// The function returns the final value of the accumulator.
// The fn function takes the current accumulator value and the current key-value pair and returns the updated accumulator value.
// The initial value of the accumulator is provided as the `initValue` argument.
//
// Example usage:
//
//	dict := NewDict[string, int]()
//	dict.Add("a", 1)
//	dict.Add("b", 2)
//	dict.Add("c", 3)
//
//	sum := dict.FoldLeft(0, func(acc int, key string, value int) int {
//	    return acc + value
//	})
//	fmt.Println(sum) // Output: 6
//
// Space Complexity:
// O(1)
//
// Time Complexity:
// O(n)
// where n is the number of key-value pairs in the Dict.
func FoldLeft[A comparable, B, C any](dict Dict[A, B], acc C, fn func(acc C, key A, value B) C) C {
	for k, v := range dict {
		acc = fn(acc, k, v)
	}
	return acc
}
