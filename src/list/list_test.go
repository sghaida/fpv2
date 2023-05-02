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
	assert.Equal(t, right, &list.List[int]{})

	var lst1 list.List[int]
	left, right = lst1.Split(5)
	assert.Equal(t, left.Size(), 0)

	lst2 := (*list.List[int])(nil)
	left, right = lst2.Split(5)
	assert.Equal(t, left.Size(), 0)
}

func TestList_Reverse(t *testing.T) {
	lst := list.FromSlice([]int{1, 2, 3, 4, 5, 6})
	reversed := lst.Reverse()
	assert.Equal(t, reversed.Head(), 6)
}

func TestList_At(t *testing.T) {
	lst := list.FromSlice([]int{0, 1, 2, 3, 4, 5})
	for index, _ := range []int{0, 1, 2, 3, 4, 5} {
		value, ok := lst.At(index)
		assert.True(t, ok)
		assert.Equal(t, value, index)
	}
	value, ok := lst.At(10)
	assert.False(t, ok)
	assert.Equal(t, value, 0)

	value, ok = lst.At(-1)
	assert.False(t, ok)
	assert.Equal(t, value, 0)
}

func TestList_AppendList(t *testing.T) {
	lst := list.FromSlice([]int{0, 1, 2, 3, 4, 5})
	appended := lst.AppendList(nil)
	assert.Equal(t, lst.Size(), 6)

	appended = lst.AppendList(list.FromSlice([]int{6, 7, 8}))
	assert.Equal(t, appended.Size(), 9)

	lst = list.FromSlice([]int{})
	appended = lst.AppendList(list.FromSlice([]int{1, 2, 3}))
	assert.Equal(t, appended.Size(), 3)
}

func TestList_Concat(t *testing.T) {
	lst := list.FromSlice([]int{0, 1, 2, 3, 4, 5})
	c1 := lst.Concat(nil)
	assert.Equal(t, c1.Size(), 6)

	c2 := lst.Concat(list.FromSlice([]int{6, 7, 8}))
	assert.Equal(t, c2.Size(), 9)

	lst = list.FromSlice([]int{})
	c3 := lst.Concat(list.FromSlice([]int{1, 2, 3}))
	assert.Equal(t, c3.Size(), 3)
}

func TestList_PrependList(t *testing.T) {
	lst := list.FromSlice([]int{0, 1, 2, 3, 4, 5})
	appended := lst.PrependList(nil)
	assert.Equal(t, lst.Size(), 6)

	appended = lst.PrependList(list.FromSlice([]int{6, 7, 8}))
	assert.Equal(t, appended.Size(), 9)

	lst = list.FromSlice([]int{})
	appended = lst.PrependList(list.FromSlice([]int{1, 2, 3}))
	assert.Equal(t, appended.Size(), 3)
}

func TestList_String(t *testing.T) {
	lst := list.FromSlice([]int{0, 1, 2, 3})
	assert.Equal(t, lst.String(), "[0, 1, 2, 3]")

	lst = list.FromSlice([]int{})
	assert.Equal(t, lst.String(), "[]")
}
