package iter

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFromMap(t *testing.T) {
	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	iter := FromMap(m)
	assert.Equal(t, iter.Size(), 3)
	assert.True(t, iter.HasNext())
	value := iter.Next()
	assert.NotEmpty(t, value)
	assert.Equal(t, iter.Count(), 2)
}

func TestMapIter_ToMap(t *testing.T) {
	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	iter := FromMap(m)
	value := iter.Next()
	toMap := iter.ToMap()
	_, ok := toMap[value.Key]
	assert.False(t, ok)
}

func TestMapIter_Clone(t *testing.T) {
	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	iter := FromMap(m)
	iter.Next()
	cloned := iter.Clone()
	assert.False(t, cloned.Size() == 3)
	assert.Equal(t, iter.Size(), cloned.Size())
}

func TestMapIter_Contains(t *testing.T) {
	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	iter := FromMap(m)
	assert.True(t, iter.Contains(1))
	assert.False(t, iter.Contains(4))
}

func TestMapIter_Filter(t *testing.T) {
	m := map[int]string{
		1: "1",
		2: "2",
		3: "3",
		4: "4",
	}
	iter := FromMap(m)
	filtered := iter.Filter(func(i int) bool {
		return i%2 == 0
	})
	assert.Equal(t, filtered.Size(), 2)
}

func TestMapIter_Reduce(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	iter := FromMap(m)
	result := iter.Reduce(func(v1, v2 int) int {
		return v1 + v2
	})
	assert.Equal(t, result, 10)
}

func TestMapIter_Fold(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	iter := FromMap(m)
	result := iter.Fold(func(acc int, key int, value int) int {
		acc += value
		return acc
	})
	assert.Equal(t, result, 10)
}

func TestMapIter_FoldLeft(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	iter := FromMap(m)
	acc := make(map[string]int)
	result := iter.FoldLeft(acc, func(key int, value int) any {
		k := strconv.Itoa(key)
		acc[k] = value
		return acc
	})
	res := result.(map[string]int)
	assert.Equal(t, len(res), 4)
	assert.Equal(t, res["1"], 1)
}

func TestMapIter_Foreach(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	iter := FromMap(m)
	iter.Foreach(func(key int, value int) {
		assert.Equal(t, key, value)
	})
}

func TestMapIter_Map(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	iter := FromMap(m)
	mapped := iter.Map(func(key int, value int) any {
		return key + value
	})
	assert.Equal(t, iter.Size(), 0)
	assert.Equal(t, mapped.Size(), 4)
}

func TestMapIter_ToSlice(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	iter := FromMap(m)
	slice := iter.ToSlice()
	assert.Equal(t, len(slice), 4)
}
