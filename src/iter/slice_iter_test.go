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
		// supposed to take 3, 4, 5, 6, 7, 8, 9
		assert.Equal(t, subIter.Size(), 7)
		subSlice := subIter.ToSlice()
		assert.Equal(t, subSlice[len(subSlice)-1], 9)
	})
}

func TestSliceIter_Filter(t *testing.T) {
	t.Run("iter of int with filter", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		iter := FromSlice(in)
		// filter even
		evenIter := iter.Filter(func(value int) bool {
			return value%2 == 0
		})
		assert.True(t, evenIter.Size() == 4)
	})

	t.Run("iter of string with filter", func(t *testing.T) {
		in := []string{"ab", "bc", "cd", "de"}
		iter := FromSlice(in)
		// filter even
		evenIter := iter.Filter(func(value string) bool {
			return value == "ab" || value == "de"
		})
		assert.True(t, evenIter.Size() == 2)
	})
}

func TestSliceIter_Map(t *testing.T) {
	t.Run("iter of int with map", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		iter := FromSlice(in)
		add1Iter := iter.Map(func(i int) any {
			return i * 10
		})
		assert.Equal(t, add1Iter.Size(), 9)
		value := add1Iter.Next()
		assert.True(t, add1Iter.HasNext())
		assert.Equal(t, value, 10)
		assert.IsType(t, 0, value)
		assert.Equal(t, add1Iter.Count(), 8)
		assert.False(t, add1Iter.HasNext())
		value = add1Iter.Next()
		assert.Equal(t, value, nil)
	})

	t.Run("iter of string with map", func(t *testing.T) {
		in := []string{"ab", "bc", "cd", "de"}
		iter := FromSlice(in)
		renameIter := iter.Map(func(value string) any {
			if value == "ab" || value == "de" {
				return "saddam"
			}
			return value
		})
		assert.Equal(t, renameIter.Size(), 4)
		value := renameIter.Next()
		assert.Equal(t, value, "saddam")
		assert.IsType(t, "", value)
		assert.Equal(t, renameIter.Count(), 3)
		assert.False(t, renameIter.HasNext())
		value = renameIter.Next()
		assert.Equal(t, value, nil)
	})
}

func TestSliceIter_Reduce(t *testing.T) {
	t.Run("iter of int with reduce", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		iter := FromSlice(in)
		result := iter.Fold(func(a, b int) int {
			return a + b
		})
		assert.Equal(t, result, 45)
	})

	t.Run("iter of string with reduce", func(t *testing.T) {
		in := []string{"a", "b", "c", "d"}
		iter := FromSlice(in)
		// filter even
		result := iter.Fold(func(a, b string) string {
			return a + b
		})
		assert.Equal(t, result, "abcd")
	})
}

func TestSliceIter_FoldLeft(t *testing.T) {
	t.Run("iter of int with FoldLeft", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		iter := FromSlice(in)
		result := iter.FoldLeft(-45, func(a any, i int) any {
			return a.(int) + i
		})
		assert.Equal(t, result, 0)
	})

	t.Run("iter of string with FoldLeft", func(t *testing.T) {
		in := []string{"a", "b", "c", "d"}
		iter := FromSlice(in)
		// filter even
		valuesMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		result := iter.FoldLeft(0, func(a any, s string) any {
			return a.(int) + valuesMap[s]
		})
		assert.Equal(t, result, 10)
	})
}

func TestSliceIter_Foreach(t *testing.T) {

	t.Run("iter of int with Foreach", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		iter := FromSlice(in)
		ch := make(chan int, 9)
		iter.Foreach(func(i int) {
			ch <- i
		})
		close(ch)
		for i := 0; i < len(in); i++ {
			number := <-ch
			assert.Equal(t, number, i+1)
		}
	})

	t.Run("iter of string with Foreach", func(t *testing.T) {
		in := []string{"a", "b", "c", "d"}
		iter := FromSlice(in)
		ch := make(chan string, 4)
		iter.Foreach(func(s string) {
			ch <- s
		})
		close(ch)
		for i := 0; i < len(in); i++ {
			number := <-ch
			assert.Equal(t, number, in[i])
		}
	})
}

func TestSliceIter_Slice(t *testing.T) {

	t.Run("iter of int with slice", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		iter := FromSlice(in)
		t.Run("proper slice", func(t *testing.T) {
			// from 4 8
			properlySliced := iter.Slice(3, 7)
			assert.Equal(t, properlySliced.Size(), 5)
			// start <= until
			untilGtStart := iter.Slice(5, 3)
			assert.Equal(t, untilGtStart.Size(), 0)
			assert.False(t, untilGtStart.HasNext())
		})

		t.Run("until greater than the size", func(t *testing.T) {
			// from 6 to 9 => size of 4
			toEnd := iter.Slice(5, 20)
			assert.Equal(t, toEnd.Size(), 4)
			assert.True(t, toEnd.HasNext())
		})

		t.Run("negative start", func(t *testing.T) {
			negativeStart := iter.Slice(-1, 4)
			assert.Equal(t, negativeStart.Size(), 0)
			assert.False(t, negativeStart.HasNext())
		})

		t.Run("exactly 1", func(t *testing.T) {
			exactly1 := iter.Slice(0, 0)
			assert.Equal(t, exactly1.Size(), 1)
			assert.True(t, exactly1.HasNext())
		})
	})
}
