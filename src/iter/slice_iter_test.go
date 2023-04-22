package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromSlice(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		iter := FromSlice(in)
		assert.Equal(t, iter.Size(), len(in))
		assert.Equal(t, iter.Count(), len(in))

		iter = FromSlice(in)
		assert.True(t, iter.HasNext())
		value := iter.Next()
		assert.Equal(t, value, 1)
		assert.Equal(t, iter.Size(), len(in)-1)

		iter.Count()
		assert.False(t, iter.HasNext())
		assert.Equal(t, iter.Next(), 0)
	})

	t.Run("str slice", func(t *testing.T) {
		in := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		iter := FromSlice(in)
		assert.Equal(t, iter.Size(), len(in))
		assert.Equal(t, iter.Count(), len(in))

		iter = FromSlice(in)
		assert.True(t, iter.HasNext())
		value := iter.Next()
		assert.Equal(t, value, "1")
		assert.Equal(t, iter.Size(), len(in)-1)

		iter.Count()
		assert.False(t, iter.HasNext())
		assert.Equal(t, iter.Next(), "")
	})
}

func TestSliceIter_Next(t *testing.T) {
	in := []int{0, 1, 2, 3, 4, 5}
	iter := FromSlice(in)
	var loc int
	for iter.HasNext() {
		value := iter.Next()
		assert.Equal(t, value, loc)
		loc++
	}
	assert.Equal(t, iter.Size(), 0)
}

func TestSliceIter_ToSlice(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		iter := FromSlice(in)
		iter.Next()
		slice := iter.ToSlice()
		assert.Equal(t, len(slice), len(in)-1)

	})

	t.Run("str slice", func(t *testing.T) {
		in := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		iter := FromSlice(in)
		iter.Next()
		slice := iter.ToSlice()
		assert.Equal(t, len(slice), len(in)-1)

	})
}

func TestSliceIter_Take(t *testing.T) {

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Run("n < iter size", func(t *testing.T) {
		iter := FromSlice(in)
		iter.Next()
		iter.Next()
		// we are @3 index 2
		subIter := iter.Take(3)
		// supposed to take 3, 4, 5
		assert.Equal(t, subIter.Size(), 3)
		subSlice := subIter.ToSlice()
		assert.Equal(t, subSlice[len(subSlice)-1], 5)
	})

	t.Run("n > iter size", func(t *testing.T) {
		iter := FromSlice(in)
		iter.Next()
		iter.Next()
		// we are @3 index 2
		subIter := iter.Take(20)
		// supposed to take 3, 4, 5, 6, 7, 8,9
		assert.Equal(t, subIter.Size(), 7)
		subSlice := subIter.ToSlice()
		assert.Equal(t, subSlice[len(subSlice)-1], 9)
	})

}
