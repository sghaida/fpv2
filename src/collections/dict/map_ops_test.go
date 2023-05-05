package dict_test

import (
	"fp/src/collections/dict"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewDict(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct[1] = "1"
	assert.Equal(t, dct[1], "1")
	assert.Equal(t, dct.Size(), 1)
}

func TestDict_Add(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct.Add(1, "1")
	assert.Equal(t, dct.Size(), 1)
}
func TestDict_Delete(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct.Add(1, "1").Add(2, "2")
	dict1 := dct.Remove(1)
	assert.Equal(t, dict1.Size(), 1)
	dict1 = dct.Remove(3)
	assert.Equal(t, dict1.Size(), 1)
}

func TestDict_Keys(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct.Add(1, "1").Add(2, "2").Add(3, "3")
	keys := dct.Keys()
	assert.Equal(t, keys.Size(), 3)
}

func TestDict_Values(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct.Add(1, "1").Add(2, "2").Add(3, "3")
	keys := dct.Values()
	assert.Equal(t, keys.Size(), 3)
}

func TestDict_KeysValues(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct.Add(1, "1").Add(2, "2").Add(3, "3")
	keys, values := dct.KeysValues()
	assert.Equal(t, keys.Size(), 3)
	assert.Equal(t, values.Size(), 3)
	assert.True(t, strconv.Itoa(keys[0]) == values[0])
}

func TestDict_Concat(t *testing.T) {
	var m1 dict.Dict[int, int]
	var m2 dict.Dict[int, int]
	m1 = dict.NewDict[int, int]().Add(1, 1).Add(2, 2)
	concatenated := m1.Concat(m2)
	assert.Equal(t, concatenated.Size(), 2)

	concatenated = m2.Concat(m1)
	assert.Equal(t, concatenated.Size(), 2)
	m2 = dict.NewDict[int, int]().Add(3, 3).Add(4, 4)
	concatenated = m1.Concat(m2)
	assert.Equal(t, concatenated.Size(), 4)
}

func TestDict_Clone(t *testing.T) {
	var m1 dict.Dict[int, int]
	m1 = dict.NewDict[int, int]().Add(1, 1).Add(2, 2)
	cloned := m1.Clone()
	assert.Equal(t, cloned.Size(), 2)
}

func TestDict_Foreach(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct.Add(1, "1").Add(2, "2").Add(3, "3")
	var acc int
	dct.Foreach(func(key int, _ string) {
		acc += key
	})
	assert.Equal(t, acc, 6)
}

func TestMap(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct.Add(1, "1").Add(2, "2").Add(3, "3")
	result := dict.Map(dct, func(key int, value string) int {
		v, _ := strconv.Atoi(value)
		return v
	})
	assert.Equal(t, result[1], 1)
}

func TestFoldLeft(t *testing.T) {
	dct := dict.NewDict[int, string]()
	dct.Add(1, "1").Add(2, "2").Add(3, "3")
	var acc int
	result := dict.FoldLeft(dct, acc, func(acc int, key int, value string) int {
		return acc + key
	})
	assert.Equal(t, result, 6)
}
