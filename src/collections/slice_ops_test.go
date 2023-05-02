package collections_test

import (
	"fp/src/collections"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestList_Size(t *testing.T) {
	lst := collections.List[int]{1, 2, 3, 4, 5}
	assert.Equal(t, lst.Size(), 5)
}

func TestList_Foreach(t *testing.T) {
	lst := collections.List[int]{1, 2, 3, 4, 5}
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
	lst := collections.List[int]{1, 2, 3, 4, 5}
	res := collections.Map(lst, func(elm int) string {
		return strconv.Itoa(elm * 2)
	})
	assert.Equal(t, res[0], "2")
}

func TestReduce(t *testing.T) {
	lst := collections.List[int]{1, 2, 3, 4, 5}
	res := collections.Reduce(lst, func(acc int, elm int) int {
		return acc + elm
	})
	assert.Equal(t, res, 15)
}

func TestFoldLeft(t *testing.T) {
	input := collections.List[int]{1, 2, 3, 4, 5}
	output := collections.FoldLeft(input, 1, func(acc int, value int) int {
		return acc * value
	})
	assert.Equal(t, output, 120)
}

func TestList_Append(t *testing.T) {
	lst := collections.List[int]{1, 2, 3, 4, 5}
	res := lst.Append(6)
	assert.Equal(t, res.Size(), 6)
	assert.Equal(t, res[5], 6)
}

func TestList_Prepend(t *testing.T) {
	lst := collections.List[int]{1, 2, 3, 4, 5}
	res := lst.Prepend(0)
	assert.Equal(t, res.Size(), 6)
	assert.Equal(t, res[0], 0)
}

func TestList_Take(t *testing.T) {
	lst := collections.List[int]{1, 2, 3, 4, 5}
	res := lst.Take(3)
	assert.Equal(t, res.Size(), 3)

	res = lst.Take(-1)
	assert.Equal(t, res.Size(), 0)

	res = lst.Take(10)
	assert.Equal(t, res.Size(), 5)
}
