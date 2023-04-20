package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("proper iter", func(t *testing.T) {
		iter, err := Range[int](0, 10, 1)
		assert.NoError(t, err)
		assert.Equal(t, iter.Size(), 10)
		next, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 9)
	})

	t.Run("end bigger than start", func(t *testing.T) {
		iter, err := Range[int](10, 0, 1)
		assert.Error(t, err)
		assert.Equal(t, iter.Size(), 0)
		next, ok := iter.Next()
		assert.False(t, ok)
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 0)
	})

	t.Run("step < 0", func(t *testing.T) {
		iter, err := Range[int](0, 10, -1)
		assert.Error(t, err)
		assert.Equal(t, iter.Size(), 0)
		next, ok := iter.Next()
		assert.False(t, ok)
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 0)
	})

	t.Run("consume iter to the ebd", func(t *testing.T) {
		iter, err := Range[int](0, 2, 1)
		assert.NoError(t, err)
		assert.Equal(t, iter.Size(), 2)

		remaining := iter.Count()
		assert.Equal(t, remaining, 2)

		next, ok := iter.Next()
		assert.False(t, ok)
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 0)
	})

}
