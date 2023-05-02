package list_test

import (
	"fp/src/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromSlice(t *testing.T) {
	lst := list.FromSlice([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, lst.Size(), 6)
	lst = lst.Append(7)
	assert.Equal(t, lst.Size(), 7)
	lst = lst.Prepend(0)
	assert.Equal(t, lst.Head(), 0)
	assert.Equal(t, lst.Tail().Head(), 1)
}

func TestList_ToSlice(t *testing.T) {
	lst := list.FromSlice([]int{1, 2, 3, 4, 5, 6})
	result := lst.ToSlice()
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, result)
}

func TestList_Split(t *testing.T) {
	lst := list.FromSlice([]int{1, 2, 3, 4, 5, 6})
	left, right := lst.Split(3)
	assert.Equal(t, left.Size(), 3)
	assert.Equal(t, left.Head(), 1)

	assert.Equal(t, right.Size(), 3)
	assert.Equal(t, right.Head(), 4)

	left, right = lst.Split(10)
	assert.Equal(t, left.Size(), 6)
	assert.Equal(t, right, (*list.List[int])(nil))

	var lst1 list.List[int]
	left, right = lst1.Split(5)
	assert.Equal(t, left.Size(), 0)

	lst2 := (*list.List[int])(nil)
	left, right = lst2.Split(5)
	assert.Equal(t, left.Size(), 0)

}
