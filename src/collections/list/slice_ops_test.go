package list_test

import (
	"fp/src/collections/list"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestList_Size(t *testing.T) {
	lst := list.List[int]{1, 2, 3, 4, 5}
	assert.Equal(t, lst.Size(), 5)
}

func TestList_Foreach(t *testing.T) {
	lst := list.List[int]{1, 2, 3, 4, 5}
	ch := make(chan int, len(lst))
	lst.Foreach(func(elm int) {
		ch <- elm
	})
	close(ch)
	var acc int
	for elm := range ch {
		acc += elm
	}
	assert.Equal(t, acc, 15)
}

func TestList_Map(t *testing.T) {
	lst := list.List[int]{1, 2, 3, 4, 5}
	res := list.Map(lst, func(elm int) string {
		return strconv.Itoa(elm * 2)
	})
	assert.Equal(t, res[0], "2")
}

func TestReduce(t *testing.T) {
	lst := list.List[int]{1, 2, 3, 4, 5}
	res := list.Reduce(lst, func(acc int, elm int) int {
		return acc + elm
	})
	assert.Equal(t, res, 15)
}

func TestFoldLeft(t *testing.T) {
	input := list.List[int]{1, 2, 3, 4, 5}
	output := list.FoldLeft(input, 1, func(acc int, value int) int {
		return acc * value
	})
	assert.Equal(t, output, 120)
}

func TestList_Append(t *testing.T) {
	lst := list.List[int]{1, 2, 3, 4, 5}
	res := lst.Append(6)
	assert.Equal(t, res.Size(), 6)
	assert.Equal(t, res[5], 6)
}

func TestList_Prepend(t *testing.T) {
	lst := list.List[int]{1, 2, 3, 4, 5}
	res := lst.Prepend(0)
	assert.Equal(t, res.Size(), 6)
	assert.Equal(t, res[0], 0)
}

func TestList_Take(t *testing.T) {
	lst := list.List[int]{1, 2, 3, 4, 5}
	res := lst.Take(3)
	assert.Equal(t, res.Size(), 3)

	res = lst.Take(-1)
	assert.Equal(t, res.Size(), 0)

	res = lst.Take(10)
	assert.Equal(t, res.Size(), 5)
}

func TestFlatten(t *testing.T) {
	lst := list.List[list.List[int]]{
		list.List[int]{1, 2},
		list.List[int]{3, 4, 5},
		list.List[int]{6},
	}
	expected := list.List[int]{1, 2, 3, 4, 5, 6}
	result := list.Flatten(lst)
	assert.Equal(t, expected, result)

	lst = list.List[list.List[int]]{
		list.List[int]{1},
		list.List[int]{},
		list.List[int]{2, 3},
	}
	expected = list.List[int]{1, 2, 3}
	result = list.Flatten(lst)
	assert.Equal(t, expected, result)

	lst = list.List[list.List[int]]{}
	expected = list.List[int](nil)
	result = list.Flatten(lst)
	assert.Equal(t, expected, result)
}

func TestFlatMap(t *testing.T) {
	lst := list.List[int]{1, 2, 3}
	mapper := func(elm int) list.List[int] {
		return list.List[int]{elm, elm * 2}
	}
	result := list.FlatMap(lst, mapper)
	assert.Equal(t, result, list.List[int]{1, 2, 2, 4, 3, 6})
}

func TestFilter(t *testing.T) {
	// Test case 1: Filter even numbers from a list of integers
	lst1 := list.List[int]{1, 2, 3, 4, 5}
	predicate1 := func(x int) bool {
		return x%2 == 0
	}
	expected1 := list.List[int]{2, 4}
	result1 := list.Filter(lst1, predicate1)
	assert.Equal(t, expected1, result1)

	// Test case 2: Filter strings that start with the letter "a" from a list of strings
	lst2 := list.List[string]{"apple", "banana", "avocado", "pear"}
	predicate2 := func(s string) bool {
		return s[0] == 'a'
	}
	expected2 := list.List[string]{"apple", "avocado"}
	result2 := list.Filter(lst2, predicate2)
	assert.Equal(t, expected2, result2)

	// Test case 3: Filter elements from an empty list
	lst3 := list.List[int]{}
	predicate3 := func(x int) bool {
		return x%2 == 0
	}
	expected3 := list.List[int](nil)
	result3 := list.Filter(lst3, predicate3)
	assert.Equal(t, expected3, result3)
}
