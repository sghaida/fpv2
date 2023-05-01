package list_test

import (
	"fp/src/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromList(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	lst := list.FromList(data)
	head := lst.Head()
	assert.Equal(t, head.Value, 1)
}

func TestList_Append(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	lst := list.FromList(data)
	res := lst.Append(7)
	assert.Equal(t, res.Head().Value, 1)
	tail := lst.Tail()
	assert.Equal(t, tail.Value, 7)
	assert.Equal(t, lst.Size(), 7)
}

func TestList_Prepend(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	lst := list.FromList(data)
	res := lst.Prepend(7)
	assert.Equal(t, res.Size(), 7)
	assert.Equal(t, res.Tail().Value, 6)
	res = list.List[int]{}
	res = res.Prepend(10)
	assert.Equal(t, res.Head().Value, 10)
}

func TestList_Delete(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	lst := list.FromList(data)
	lst.Delete(1)
	assert.Equal(t, lst.Head().Value, 2)
	assert.Equal(t, lst.Size(), 5)

	res := lst.Delete(6)
	assert.Equal(t, res.Size(), 4)

	lst = &list.List[int]{}
	lst.Delete(5)
	assert.Equal(t, lst.Size(), 0)
}

func TestList_Find(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	lst := list.FromList(data)

	elem := lst.Find(3)
	assert.Equal(t, elem.Value, 3)

	elem = lst.Find(1)
	assert.Equal(t, elem.Value, 1)

	elem = lst.Find(10)
	assert.Equal(t, elem, list.Node[int]{})
}
