package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("proper iter", func(t *testing.T) {
		iter, err := Range[int](0, 9, 1)
		assert.NoError(t, err)
		assert.Equal(t, iter.Size(), 10)
		ok := iter.HasNext()
		assert.True(t, ok)
		next := iter.Next()
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 9)
	})

	t.Run("end bigger than start", func(t *testing.T) {
		iter, err := Range[int](10, 0, 1)
		assert.Error(t, err)
		assert.Equal(t, iter.Size(), 0)
		ok := iter.HasNext()
		assert.False(t, ok)
		next := iter.Next()
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 0)
	})

	t.Run("step < 0", func(t *testing.T) {
		iter, err := Range[int](0, 10, -1)
		assert.Error(t, err)
		assert.Equal(t, iter.Size(), 0)
		ok := iter.HasNext()
		assert.False(t, ok)
		next := iter.Next()
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 0)
	})

	t.Run("consume iter to the end", func(t *testing.T) {
		iter, err := Range[int](0, 2, 1)
		assert.NoError(t, err)
		assert.Equal(t, iter.Size(), 3)

		remaining := iter.Count()
		assert.Equal(t, remaining, 3)

		ok := iter.HasNext()
		assert.False(t, ok)
		next := iter.Next()
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 0)
		assert.Equal(t, iter.Size(), 0)
	})

}

func TestRangeIter_Take(t *testing.T) {
	t.Run("start + n <= end", func(t *testing.T) {
		iter, _ := Range[int](0, 5, 1)
		iter.Next()
		remaining := iter.Take(4, 1)
		assert.Equal(t, remaining.Size(), 4)
		value := remaining.Next()
		assert.Equal(t, value, 1)
		assert.Equal(t, remaining.Size(), 3)
	})

	t.Run("start +n > end", func(t *testing.T) {
		iter, _ := Range[int](0, 4, 1)
		iter.Next()
		remaining := iter.Take(7, 1)
		assert.Equal(t, remaining.Size(), 4)
		value := iter.Next()
		assert.Equal(t, value, 1)
		assert.Equal(t, iter.Size(), 3)
		//consume to end
		remaining.Next()
		remaining.Next()
		remaining.Next()
		remaining.Next()
		assert.Equal(t, remaining.Size(), 0)
	})
}
