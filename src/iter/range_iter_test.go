package iter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("proper iter", func(t *testing.T) {
		iter, err := Range[int](0, 10, 1)
		assert.NoError(t, err)
		assert.Equal(t, iter.Size(), 11)
		ok := iter.HasNext()
		assert.True(t, ok)
		next := iter.Next()
		assert.Equal(t, next, 0)
		assert.Equal(t, iter.Count(), 10)
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

		count := iter.Count()
		assert.Equal(t, count, 3)

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

func TestRangeIter_Filter(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	// filter even
	evenIter := iter.Filter(func(value int) bool {
		return value%2 == 0
	})
	assert.True(t, evenIter.Size() == 4)
}

func TestRangeIter_Map(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	mappedIter := iter.Map(func(i int) any {
		return fmt.Sprintf("%d", i)
	})
	assert.Equal(t, mappedIter.Next(), "1")
}

func TestRangeIter_Reduce(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	value := iter.Reduce(func(i int, i2 int) int {
		return i + i2
	})
	assert.Equal(t, value, 45)
}

func TestRangeIter_Fold(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	value := iter.Fold(func(i int, i2 int) int {
		return i + i2
	})
	assert.Equal(t, value, 45)
}

func TestRangeIter_FoldLeft(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	value := iter.FoldLeft("", func(i any, i2 int) any {
		return i.(string) + strconv.Itoa(i2)
	})
	assert.Equal(t, value, "123456789")
}

func TestRangeIter_Foreach(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	ch := make(chan int, 9)
	iter.Foreach(func(i int) {
		ch <- i
	})
	close(ch)
	for i := 1; i < 9; i++ {
		value := <-ch
		assert.Equal(t, value, i)
	}
}

func TestRangeIter_Slice(t *testing.T) {
	t.Run("proper slice", func(t *testing.T) {
		iter, _ := Range[int](1, 9, 1)
		slicedIter := iter.Slice(4, 7)
		assert.Equal(t, slicedIter.Size(), 4)
	})

	t.Run("from is greater than until", func(t *testing.T) {
		iter, _ := Range[int](1, 9, 1)
		slicedIter := iter.Slice(4, 1)
		assert.Equal(t, slicedIter.Size(), 0)
	})

	t.Run("from is greater than iter size", func(t *testing.T) {
		iter, _ := Range[int](1, 9, 1)
		slicedIter := iter.Slice(15, 20)
		assert.Equal(t, slicedIter.Size(), 0)
	})

	t.Run("from is less than zero", func(t *testing.T) {
		iter, _ := Range[int](1, 9, 1)
		slicedIter := iter.Slice(-1, 20)
		assert.Equal(t, slicedIter.Size(), 0)
	})

}

func TestRangeIter_Clone(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	cloned := iter.Clone()
	assert.Equal(t, iter, cloned)
	cloned.Next()
	assert.NotEqual(t, iter, cloned)
}

func TestRangeIter_Drop(t *testing.T) {
	t.Run("proper drop", func(t *testing.T) {
		iter, _ := Range[int](1, 9, 1)
		dropped := iter.Drop(4)
		assert.Equal(t, dropped.Size(), 5)
	})

	t.Run("> size", func(t *testing.T) {
		iter, _ := Range[int](1, 9, 1)
		dropped := iter.Drop(10)
		assert.Equal(t, dropped.Size(), 0)
	})

}

func TestRangeIter_Contains(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	status := iter.Contains(10)
	assert.False(t, status)

	iter, _ = Range[int](1, 9, 1)
	status = iter.Contains(5)
	assert.True(t, status)
}

func TestRangeIter_ToIter(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	pureIter := iter.ToIter()
	assert.Equal(t, pureIter.Size(), 9)
}

func TestRangeIter_ToSlice(t *testing.T) {
	iter, _ := Range[int](1, 9, 1)
	slice := iter.ToSlice()
	assert.Equal(t, slice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
