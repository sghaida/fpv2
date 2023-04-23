package src

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeft(t *testing.T) {
	left := Left[error, int](errors.New("le"))

	assert.True(t, left.IsLeft())
	assert.False(t, left.IsRight())
	assert.IsType(t, Either[error, int]{}, left)

	lv, err := left.TakeLeft()
	assert.NoError(t, err)
	assert.IsType(t, errors.New("le"), lv)
	assert.Equal(t, lv, errors.New("le"))
	rv, err := left.TakeRight()
	assert.Equal(t, rv, 0)
	assert.Equal(t, err, ErrorLeftValue)
	assert.IsType(t, 0, rv)
}

func TestRight(t *testing.T) {
	t.Run("primitive value", func(t *testing.T) {
		right1 := Right[error, int](10)

		// right1 assertions
		assert.True(t, right1.IsRight())
		assert.False(t, right1.IsLeft())
		assert.IsType(t, Either[error, int]{}, right1)

		rv1, err := right1.TakeRight()
		assert.Equal(t, rv1, 10)
		assert.Equal(t, err, nil)
		assert.IsType(t, 0, rv1)

		// left assertions
		lv1, err := right1.TakeLeft()
		assert.Error(t, err)
		assert.IsType(t, ErrorRightValue, err)
		assert.Equal(t, lv1, nil)
	})

	t.Run("nil array", func(t *testing.T) {
		// this should fail as we are passing unreferenced array
		right2 := Right[error, []int](nil)
		assert.True(t, right2.IsLeft())
		assert.False(t, right2.IsRight())

		rv2, err := right2.TakeRight()

		assert.Equal(t, err, ErrorLeftValue)
		assert.IsType(t, []int{}, rv2)
	})

	t.Run("proper array", func(t *testing.T) {
		right3 := Right[error, []int]([]int{1, 2, 3})
		assert.True(t, right3.IsRight())
		assert.False(t, right3.IsLeft())

		rv3, err := right3.TakeRight()

		assert.Equal(t, err, nil)
		assert.Equal(t, rv3, []int{1, 2, 3})
		assert.IsType(t, []int{}, rv3)
	})
}

func TestEither_Unwrap(t *testing.T) {
	l := Left[error, string](errors.New("some error"))
	lv, rv := l.Unwrap()

	assert.Equal(t, rv, errors.New("some error"))
	assert.Equal(t, lv, "")

	r := Right[error, string]("saddam")
	lv, rv = r.Unwrap()

	assert.Equal(t, lv, "saddam")
	assert.Equal(t, rv, nil)
}

func TestEither_Swap(t *testing.T) {

	t.Run("swap left", func(t *testing.T) {
		l := Left[error, string](errors.New("some error"))
		lw := l.Swap()
		assert.True(t, lw.IsRight())

		assert.IsType(t, "", lw.left)
		assert.Equal(t, "", lw.left)

		assert.IsType(t, errors.New(""), lw.right)
		assert.Equal(t, lw.right, errors.New("some error"))
	})

	t.Run("swap right", func(t *testing.T) {
		r := Right[error, string]("saddam")
		rw := r.Swap()
		assert.True(t, rw.IsLeft())

		assert.IsType(t, "", rw.left)
		assert.Equal(t, "saddam", rw.left)

		assert.IsType(t, nil, rw.right)
		assert.Equal(t, rw.right, nil)
	})
}

func TestEither_Take(t *testing.T) {
	t.Run("primitive value", func(t *testing.T) {
		right1 := Right[error, int](10)

		// right1 assertions
		assert.True(t, right1.IsRight())
		assert.False(t, right1.IsLeft())
		assert.IsType(t, Either[error, int]{}, right1)

		rv1, err := right1.Take()
		assert.Equal(t, rv1, 10)
		assert.Equal(t, err, nil)
		assert.IsType(t, 0, rv1)
	})

	t.Run("nil array", func(t *testing.T) {
		// this should fail as we are passing unreferenced array
		right2 := Right[error, []int](nil)
		assert.True(t, right2.IsLeft())
		assert.False(t, right2.IsRight())

		rv2, err := right2.Take()

		assert.Equal(t, err, ErrorLeftValue)
		assert.IsType(t, []int{}, rv2)
	})

	t.Run("proper array", func(t *testing.T) {
		right3 := Right[error, []int]([]int{1, 2, 3})
		assert.True(t, right3.IsRight())
		assert.False(t, right3.IsLeft())

		rv3, err := right3.Take()

		assert.Equal(t, err, nil)
		assert.Equal(t, rv3, []int{1, 2, 3})
		assert.IsType(t, []int{}, rv3)
	})
}

func TestEither_TakeOr(t *testing.T) {
	t.Run("valid right", func(t *testing.T) {
		r := Right[error, int](10)
		value := r.TakeOr(20)
		assert.Equal(t, value, 10)
	})

	t.Run("invalid right", func(t *testing.T) {
		l := Left[error, int](errors.New("missing value"))
		value := l.TakeOr(10)
		assert.Equal(t, value, 10)
	})
}

func TestEither_TakeOrElse(t *testing.T) {
	t.Run("valid right", func(t *testing.T) {
		r := Right[error, int](10)
		value := r.TakeOrElse(func() int {
			return 20
		})
		assert.Equal(t, value, 10)
	})

	t.Run("invalid right", func(t *testing.T) {
		l := Left[error, int](errors.New("missing value"))
		value := l.TakeOrElse(func() int {
			return 10
		})
		assert.Equal(t, value, 10)
	})
}

func TestEither_Or(t *testing.T) {
	t.Run("valid right", func(t *testing.T) {
		r := Right[error, int](10)
		value := r.Or(Right[error, int](20))
		assert.Equal(t, value.right, 10)
	})

	t.Run("invalid right", func(t *testing.T) {
		l := Left[error, int](errors.New("missing value"))
		value := l.Or(Right[error, int](10))
		assert.Equal(t, value.right, 10)
	})
}

func TestEither_OrElse(t *testing.T) {
	t.Run("valid right", func(t *testing.T) {
		r := Right[error, int](10)
		value := r.OrElse(func() Either[error, int] {
			return Right[error, int](20)
		})
		assert.Equal(t, value.right, 10)
	})

	t.Run("invalid right", func(t *testing.T) {
		l := Left[error, int](errors.New("missing value"))
		value := l.OrElse(func() Either[error, int] {
			return Right[error, int](10)
		})
		assert.Equal(t, value.right, 10)
	})
}

func TestEither_Map(t *testing.T) {
	t.Run("valid right", func(t *testing.T) {
		r := Right[error, int](10)
		value := r.Map(func(i int) any {
			return i + 10
		})
		assert.True(t, value.IsRight())
		assert.Equal(t, value.right, 20)
	})

	t.Run("invalid right", func(t *testing.T) {
		l := Left[error, int](errors.New("missing value"))
		value := l.Map(func(i int) any {
			return Right[error, int](10)
		})
		assert.True(t, value.IsLeft())
		assert.Equal(t, value.left, errors.New("missing value"))
		assert.IsType(t, errors.New("missing value"), value.left)
		assert.Equal(t, value.right, nil)
	})
}

func TestEither_Fold(t *testing.T) {
	t.Run("valid right", func(t *testing.T) {
		r := Right[error, int](10)
		value := r.Fold(
			func(err error) any {
				return errors.New("this should not happen")
			},
			func(i int) any {
				return i + 10
			},
		)
		switch v := value.(type) {
		case int:
			assert.Equal(t, v, 20)
		case error:
			assert.Fail(t, "expected err, got int")
		}
	})

	t.Run("invalid right", func(t *testing.T) {
		l := Left[error, int](errors.New("missing value"))
		value := l.Fold(
			func(err error) any {
				return errors.New("not applicable fold")
			},
			func(i int) any {
				return i + 10
			},
		)
		switch v := value.(type) {
		case int:
			assert.Fail(t, "expected err, got int")
		case error:
			assert.Equal(t, v, errors.New("not applicable fold"))
		}
	})
}

func TestEither_Exists(t *testing.T) {

	t.Run("exists", func(t *testing.T) {
		r := Right[error, int](10)
		status := r.Exists(func(value int) bool {
			return value == 10
		})
		assert.True(t, status)
	})

	t.Run("doesn't exist", func(t *testing.T) {
		r := Right[error, int](10)
		status := r.Exists(func(value int) bool {
			return value == 20
		})
		assert.False(t, status)

		l := Left[error, int](errors.New("value is not presented"))
		status = l.Exists(func(value int) bool {
			return value == 20
		})
		assert.False(t, status)
	})
}

func TestEither_FlatMap(t *testing.T) {

	t.Run("right side 1 level", func(t *testing.T) {
		r := Right[error, int](10)
		fr := r.FlatMap(func(value int) Either[error, any] {
			return Right[error, any](value + 10)
		})
		assert.True(t, fr.IsRight())
		assert.Equal(t, fr.right, 20)
		assert.Equal(t, fr.left, nil)
	})

	t.Run("left side 1 level", func(t *testing.T) {
		r := Left[error, int](errors.New("missing value"))
		fr := r.FlatMap(func(value int) Either[error, any] {
			return Right[error, any](10 + 10)
		})
		assert.True(t, fr.IsLeft())
		assert.Equal(t, fr.right, nil)
		assert.Equal(t, fr.left, errors.New("missing value"))
	})

	t.Run("right side 2 level with flatten", func(t *testing.T) {
		innerR := Right[error, int](10)
		r := Right[error, Either[error, int]](innerR)
		fr := r.FlatMap(func(value Either[error, int]) Either[error, any] {
			if value.IsRight() {
				v, _ := value.TakeRight()
				return Right[error, any](v + 10)
			}
			return any(value).(Either[error, any])
		})
		assert.True(t, fr.IsRight())
		assert.Equal(t, fr.right, 20)
		assert.Equal(t, fr.left, nil)
	})

	t.Run("left side 2 level with flatten", func(t *testing.T) {
		r := Left[error, int](errors.New("missing value"))
		rl := Right[error, Either[error, int]](r)
		fr := rl.FlatMap(func(value Either[error, int]) Either[error, any] {
			if value.IsRight() {
				v, _ := value.TakeRight()
				return Right[error, any](v + 10)
			}
			left, _ := value.TakeLeft()
			return Left[error, any](left)
		})
		assert.True(t, fr.IsLeft())
		assert.Equal(t, fr.right, nil)
		assert.Equal(t, fr.left, errors.New("missing value"))
	})
}

func TestEither_ToOption(t *testing.T) {
	t.Run("some", func(t *testing.T) {
		r := Right[error, int](10)
		some := r.ToOption()
		assert.Equal(t, some.value, 10)
		assert.IsType(t, 10, some.value)
	})

	t.Run("none", func(t *testing.T) {
		r := Left[error, int](errors.New("missing value"))
		none := r.ToOption()
		assert.Equal(t, none.value, 0)
		assert.IsType(t, 10, none.value)
	})
}
