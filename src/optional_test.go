package src

import (
	"crypto/sha1"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewOptional(t *testing.T) {
	var emptySlice []string
	var nullVariable *string
	type optionTestCase[A any] struct {
		name         string
		expectedType Type
		expectsError bool
		value        A
	}
	tt := []optionTestCase[interface{}]{
		{
			name:         "Optional Number",
			value:        10,
			expectedType: SomeType,
		}, {
			name:         "Optional string",
			value:        "some option",
			expectedType: SomeType,
		}, {
			name:         "optional array",
			value:        []int{1, 2, 3, 4},
			expectedType: SomeType,
		}, {
			name:         "optional map",
			value:        map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			expectedType: SomeType,
		}, {
			name:         "Optional nil value",
			value:        nil,
			expectedType: NoneType,
		}, {
			name:         "optional struct",
			expectedType: SomeType,
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
		}, {
			name:         "optional pointer type",
			value:        nullVariable,
			expectedType: NoneType,
		}, {
			name:         "emptySlice",
			expectedType: NoneType,
			value:        emptySlice,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			if option.t == SomeType {
				assert.Equal(t, option.Unwrap(), tc.value)
			}
			assert.Equal(t, option.t, tc.expectedType)
		})
	}
}

func TestNone(t *testing.T) {
	option := None[int]()
	assert.Equal(t, option.t, NoneType)
}

func TestSome(t *testing.T) {
	var emptySlice []string
	var nullVariable *string
	type optionTestCase[A any] struct {
		name         string
		expectedType Type
		expectsError bool
		value        A
	}
	tt := []optionTestCase[interface{}]{
		{
			name:         "Some Number",
			value:        10,
			expectedType: SomeType,
		}, {
			name:         "Some string",
			value:        "some option",
			expectedType: SomeType,
		}, {
			name:         "Some array",
			value:        []int{1, 2, 3, 4},
			expectedType: SomeType,
		}, {
			name:         "Some map",
			value:        map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			expectedType: SomeType,
		}, {
			name:         "Some nil value",
			value:        nil,
			expectedType: NoneType,
			expectsError: true,
		}, {
			name:         "Some struct",
			expectedType: SomeType,
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
		}, {
			name:         "Some emptyString",
			value:        nullVariable,
			expectsError: true,
			expectedType: NoneType,
		}, {
			name:         "Some nilSlice",
			expectedType: NoneType,
			expectsError: true,
			value:        emptySlice,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option, err := Some(tc.value)
			if err != nil {
				assert.Equal(t, tc.expectsError, true)
				assert.Equal(t, option.IsNone(), true)
				assert.Equal(t, option.GetType(), NoneType)
				return
			}
			assert.Equal(t, option.IsSome(), true)
			assert.Equal(t, option.Get(), tc.value)
			assert.Equal(t, option.t, tc.expectedType)
			assert.Equal(t, option.GetType(), SomeType)
		})
	}
}

func TestOption_Apply(t *testing.T) {
	//var emptySlice []string
	//var nullVariable *string
	type testData[A any] struct {
		name          string
		fn            func(value A) A
		value         A
		expectedValue A
		expectedType  Type
	}
	tt := []testData[any]{
		{
			name: "apply on number",
			fn: func(value any) any {
				return value.(int) + 1
			},
			value:         10,
			expectedType:  SomeType,
			expectedValue: 11,
		}, {
			name: "apply to string",
			fn: func(value any) any {
				return fmt.Sprintf("%v abu ghaida", value)
			},
			value:         "saddam",
			expectedValue: "saddam abu ghaida",
			expectedType:  SomeType,
		}, {
			name: "apply on bool",
			fn: func(_ any) any {
				return true
			},
			value:         false,
			expectedValue: true,
			expectedType:  SomeType,
		}, {
			name: "apply on none type",
			fn: func(value any) any {
				return "nil"
			},
			value:         nil,
			expectedValue: nil,
			expectedType:  NoneType,
		}, {
			name: "apply on empty slice",
			fn: func(value any) any {
				return []string{"saddam"}
			},
			value:         []string(nil),
			expectedValue: []string(nil),
			expectedType:  NoneType,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			if option.GetType() == SomeType {
				assert.Equal(t, option.Get(), tc.value)
			}
			assert.Equal(t, option.GetType(), tc.expectedType)
			// apply
			option = option.Apply(tc.fn)
			if option.GetType() == SomeType {
				assert.Equal(t, tc.expectedValue, option.Get())
			}
			assert.Equal(t, option.GetType(), tc.expectedType)
		})
	}
}

func TestOption_Take(t *testing.T) {
	var emptySlice []string
	var nullVariable *string
	type optionTestCase[A any] struct {
		name         string
		expectedType Type
		expectsError bool
		value        A
	}
	tt := []optionTestCase[interface{}]{
		{
			name:         "Some Number",
			value:        10,
			expectedType: SomeType,
		}, {
			name:         "Some string",
			value:        "some option",
			expectedType: SomeType,
		}, {
			name:         "Some array",
			value:        []int{1, 2, 3, 4},
			expectedType: SomeType,
		}, {
			name:         "Some map",
			value:        map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			expectedType: SomeType,
		}, {
			name:         "Some nil value",
			value:        nil,
			expectsError: true,
			expectedType: NoneType,
		}, {
			name:         "Some struct",
			expectedType: SomeType,
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
		}, {
			name:         "Some pointer type",
			value:        nullVariable,
			expectsError: true,
			expectedType: NoneType,
		}, {
			name:         "Some emptySlice",
			expectedType: NoneType,
			expectsError: true,
			value:        emptySlice,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option, _ := Some(tc.value)
			value, err := option.Take()
			if err != nil {
				assert.Equal(t, tc.expectsError, true)
				return
			}
			assert.Equal(t, value, tc.value)
			assert.False(t, tc.expectsError)

		})
	}
}

func TestOption_TakeOr(t *testing.T) {

	type testData[A any] struct {
		name       string
		alt        A
		value      A
		isNoneType bool
	}

	var emptySlice []string
	emptySliceAlt := []string{"some alternative slice as ref"}

	var nullVariable *string
	nullVariableAlt := "some alternative string as ref"
	tt := []testData[interface{}]{
		{
			name:  "Some Number",
			value: 10,
			alt:   11,
		}, {
			name:  "Some string",
			value: "some option",
			alt:   "some other option",
		}, {
			name:  "Some array",
			value: []int{1, 2, 3, 4},
			alt:   []int{5, 6, 7, 8},
		}, {
			name:  "Some map",
			value: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			alt:   map[int]int{5: 5, 6: 6, 7: 7, 8: 8},
		}, {
			name: "Some struct",
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			alt: struct {
				name string
				age  int
			}{
				name: "saddam1",
				age:  222,
			},
		}, {
			name:       "Some pointer type",
			value:      nullVariable,
			alt:        &nullVariableAlt,
			isNoneType: true,
		}, {
			name:       "Some emptySlice",
			value:      emptySlice,
			alt:        &emptySliceAlt,
			isNoneType: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			value := option.TakeOr(tc.alt)
			if tc.isNoneType {
				assert.Equal(t, value, tc.alt)
				return
			}
			assert.Equal(t, value, tc.value)
		})
	}
}

func TestOption_TakeOrElse(t *testing.T) {

	type testData[A any] struct {
		name       string
		fn         func() A
		value      A
		isNoneType bool
	}

	var emptySlice []string
	emptySliceAlt := []string{"some alternative slice as ref"}

	var nullVariable *string
	nullVariableAlt := "some alternative string as ref"
	tt := []testData[interface{}]{
		{
			name:  "Some Number",
			value: 10,
			fn: func() any {
				return 11
			},
		}, {
			name:  "Some string",
			value: "some option",
			fn: func() any {
				return "some other option"
			},
		}, {
			name:  "Some array",
			value: []int{1, 2, 3, 4},
			fn: func() any {
				return []int{5, 6, 7, 8}
			},
		}, {
			name:  "Some map",
			value: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			fn: func() any {
				return map[int]int{5: 5, 6: 6, 7: 7, 8: 8}
			},
		}, {
			name: "Some struct",
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			fn: func() any {
				return struct {
					name string
					age  int
				}{
					name: "saddam1",
					age:  222,
				}
			},
		}, {
			name:  "Some pointer type",
			value: nullVariable,
			fn: func() any {
				return &nullVariableAlt
			},
			isNoneType: true,
		}, {
			name:  "Some emptySlice",
			value: emptySlice,
			fn: func() any {
				return &emptySliceAlt
			},
			isNoneType: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			value := option.TakeOrElse(tc.fn)
			if tc.isNoneType {
				assert.Equal(t, value, tc.fn())
				return
			}
			assert.Equal(t, value, tc.value)
		})
	}
}

func TestOption_Or(t *testing.T) {
	type testData[A any] struct {
		name          string
		value         A
		alt           Option[A]
		expectedValue A
		isNoneType    bool
	}

	var emptySlice []string
	emptySliceAlt := []string{"some alternative slice"}

	var emptyStr = ""
	emptyVariableAlt := "some alternative string"
	tt := []testData[any]{
		{
			name:          "Some Number",
			value:         10,
			expectedValue: 10,
			alt:           NewOptional[any](11),
		}, {
			name:          "Some string",
			value:         "some option",
			expectedValue: "some option",
			alt:           NewOptional[any]("some other option"),
		}, {
			name:          "Some array",
			value:         []int{1, 2, 3, 4},
			expectedValue: []int{1, 2, 3, 4},
			alt:           NewOptional[any]([]int{5, 6, 7, 8}),
		}, {
			name:          "Some map",
			value:         map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			expectedValue: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			alt:           NewOptional[any](map[int]int{5: 5, 6: 6, 7: 7, 8: 8}),
		}, {
			name: "Some struct",
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			expectedValue: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			alt: NewOptional[any](struct {
				name string
				age  int
			}{
				name: "saddam1",
				age:  222,
			}),
		}, {
			name:          "Some empty string",
			value:         emptyStr,
			expectedValue: emptyStr,
			alt:           NewOptional[any](emptyVariableAlt),
		}, {
			name:          "Some emptySlice",
			value:         emptySlice,
			expectedValue: emptySliceAlt,
			alt:           NewOptional[any](emptySliceAlt),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			altOptValue := option.Or(tc.alt)
			assert.Equal(t, altOptValue.value, tc.expectedValue)
		})
	}
}

func TestOption_OrElse(t *testing.T) {
	type testData[A any] struct {
		name          string
		value         A
		fn            func() Option[A]
		expectedValue A
		isNoneType    bool
	}

	var emptySlice []string
	emptySliceAlt := []string{"some alternative slice"}

	var emptyStr = ""
	emptyVariableAlt := "some alternative string"
	tt := []testData[any]{
		{
			name:          "Some Number",
			value:         10,
			expectedValue: 10,
			fn: func() Option[any] {
				return NewOptional[any](11)
			},
		}, {
			name:          "Some string",
			value:         "some option",
			expectedValue: "some option",
			fn: func() Option[any] {
				return NewOptional[any]("some other option")
			},
		}, {
			name:          "Some array",
			value:         []int{1, 2, 3, 4},
			expectedValue: []int{1, 2, 3, 4},
			fn: func() Option[any] {
				return NewOptional[any]([]int{5, 6, 7, 8})
			},
		}, {
			name:          "Some map",
			value:         map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			expectedValue: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			fn: func() Option[any] {
				return NewOptional[any](map[int]int{5: 5, 6: 6, 7: 7, 8: 8})
			},
		}, {
			name: "Some struct",
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			expectedValue: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			fn: func() Option[any] {
				return NewOptional[any](struct {
					name string
					age  int
				}{
					name: "saddam1",
					age:  222,
				})
			},
		}, {
			name:          "Some empty string",
			value:         emptyStr,
			expectedValue: emptyStr,
			fn: func() Option[any] {
				return NewOptional[any](emptyVariableAlt)
			},
		}, {
			name:          "Some emptySlice",
			value:         emptySlice,
			expectedValue: emptySliceAlt,
			fn: func() Option[any] {
				return NewOptional[any](emptySliceAlt)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			altOptValue := option.OrElse(tc.fn)
			assert.Equal(t, altOptValue.value, tc.expectedValue)
		})
	}
}

func TestOption_IfSome(t *testing.T) {
	type testData[A any] struct {
		name       string
		value      A
		fn         func(value A) A
		isNoneType bool
	}

	var emptySlice []string
	emptySliceAlt := []string{"some alternative slice"}

	var emptyStr = ""
	emptyVariableAlt := "some alternative string"
	tt := []testData[any]{
		{
			name:  "Some Number",
			value: 10,
			fn: func(value any) any {
				return value.(int) + 1
			},
		}, {
			name:  "Some string",
			value: "saddam",
			fn: func(value any) any {
				return fmt.Sprintf("%s abu ghaida", value)
			},
		}, {
			name:  "Some array",
			value: []int{1, 2, 3, 4},
			fn: func(value any) any {
				return append(value.([]int), 5, 6, 7, 8)
			},
		}, {
			name:  "Some map",
			value: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			fn: func(value any) any {
				v := value.(map[int]int)
				v[5] = 5
				v[6] = 6
				return v
			},
		}, {
			name: "Some struct",
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			fn: func(value any) any {
				v := value.(struct {
					name string
					age  int
				})
				v.name = "saddam1"
				v.age = 222
				return v
			},
		}, {
			name:  "Some empty string",
			value: emptyStr,
			fn: func(_ any) any {
				return emptyVariableAlt
			},
		}, {
			name:  "Some emptySlice",
			value: emptySlice,
			fn: func(_ any) any {
				return emptySliceAlt
			},
			isNoneType: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			altOptValue, err := option.IfSome(tc.fn)
			if err != nil {
				assert.True(t, tc.isNoneType)
				return
			}
			assert.Equal(t, altOptValue, tc.fn(option.Unwrap()))
		})
	}
}

func TestOption_IfNone(t *testing.T) {
	type testData[A any] struct {
		name       string
		value      A
		fn         func() A
		isNoneType bool
	}

	var emptySlice []string
	emptySliceAlt := []string{"some alternative slice"}

	var emptyStr = ""
	emptyVariableAlt := "some alternative string"
	tt := []testData[any]{
		{
			name:  "Some Number",
			value: 10,
			fn: func() any {
				return 11
			},
		}, {
			name:  "Some string",
			value: "some option",
			fn: func() any {
				return "some other option"
			},
		}, {
			name:  "Some array",
			value: []int{1, 2, 3, 4},
			fn: func() any {
				return []int{5, 6, 7, 8}
			},
		}, {
			name:  "Some map",
			value: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			fn: func() any {
				return map[int]int{5: 5, 6: 6, 7: 7, 8: 8}
			},
		}, {
			name: "Some struct",
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			fn: func() any {
				return struct {
					name string
					age  int
				}{
					name: "saddam1",
					age:  222,
				}
			},
		}, {
			name:  "Some empty string",
			value: emptyStr,
			fn: func() any {
				return emptyVariableAlt
			},
		}, {
			name:  "Some emptySlice",
			value: emptySlice,
			fn: func() any {
				return emptySliceAlt
			},
			isNoneType: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			altOptValue, err := option.IfNone(tc.fn)
			if err != nil {
				assert.False(t, tc.isNoneType)
				return
			}
			assert.Equal(t, altOptValue, tc.fn())
		})
	}
}

func TestOption_String(t *testing.T) {
	type optionTestCase[A any] struct {
		name        string
		expectedStr string
		value       A
	}
	tt := []optionTestCase[interface{}]{
		{
			name:        "Optional Number",
			value:       10,
			expectedStr: "Some[10]",
		}, {
			name:        "Optional string",
			value:       "some option",
			expectedStr: "Some[some option]",
		}, {
			name:        "Optional nil value",
			value:       nil,
			expectedStr: "None",
		}, {
			name:        "optional array",
			value:       []int{1, 2, 3, 4},
			expectedStr: "Some[[1 2 3 4]]",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			assert.Equal(t, option.String(), tc.expectedStr)
		})
	}
}

func TestOption_Map(t *testing.T) {

	type testData[A, B any] struct {
		name  string
		value A
		fn    Mapper[A, B]
	}

	var emptySlice []string
	emptySliceAlt := []string{"some alternative slice"}

	var emptyStr = ""
	emptyVariableAlt := "some alternative string"
	tt := []testData[any, any]{
		{
			name:  "Some Number to string",
			value: 10,
			fn: func(value any) any {
				return fmt.Sprintf("%v", value)
			},
		}, {
			name:  "Some string to int",
			value: "saddam",
			fn: func(value any) any {
				str := fmt.Sprintf("%s abu ghaida", value)
				hash := sha1.New()
				hash.Write([]byte(str))
				return hash.Size()
			},
		}, {
			name:  "Some array to map",
			value: []int{1, 2, 3, 4},
			fn: func(value any) any {
				m := map[int]int{}
				for i, value := range value.([]int) {
					m[i+1] = value
				}
				return m
			},
		}, {
			name:  "Some map to array",
			value: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			fn: func(value any) any {
				v := value.(map[int]int)
				var arr []int
				for i := 0; i < 5; i++ {
					arr = append(arr, v[i])
				}

				return arr
			},
		}, {
			name: "Some struct map on name value",
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			fn: func(value any) any {
				v := value.(struct {
					name string
					age  int
				})

				return v.name
			},
		}, {
			name:  "Some empty string",
			value: emptyStr,
			fn: func(_ any) any {
				return emptyVariableAlt
			},
		}, {
			name:  "Some emptySlice",
			value: emptySlice,
			fn: func(_ any) any {
				return emptySliceAlt
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			app := option.Map(tc.fn)
			if option.IsNone() {
				assert.True(t, app.IsNone())
				return
			}
			assert.Equal(t, option.t, app.t)
			assert.Equal(t, app.Get(), tc.fn(option.Unwrap()))
		})
	}
}

func TestOption_Flatten(t *testing.T) {

	type testData[A Option[any]] struct {
		name     string
		value    A
		expected Option[any]
	}

	tt := []testData[Option[any]]{
		{
			name:     "Some Option of Option of Int",
			value:    NewOptional[any](NewOptional[any](10)),
			expected: NewOptional[any](10),
		}, {
			name:     "Some Option of Int",
			value:    NewOptional[any](10),
			expected: NewOptional[any](10),
		},
		{
			name:     "Some Option of Option of string",
			value:    NewOptional[any](NewOptional[any]("saddam")),
			expected: NewOptional[any]("saddam"),
		}, {
			name:     "Some Option of string",
			value:    NewOptional[any]("saddam"),
			expected: NewOptional[any]("saddam"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			flattened := tc.value.Flatten()
			assert.Equal(t, flattened.Get(), tc.expected.Get())

		})
	}
}

func TestOption_FlatMap(t *testing.T) {
	type testData[A any, B Option[any]] struct {
		name         string
		value        A
		fn           func(A any) B
		expected     Option[any]
		expectsError bool
	}

	tt := []testData[any, Option[any]]{
		{
			name:         "Some Option of Int",
			value:        NewOptional[any](10),
			expected:     NewOptional[any]("11"),
			expectsError: false,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		}, {
			name:         "Some Option of Option of Int should return option of string",
			value:        NewOptional[any](NewOptional[any](10)),
			expected:     NewOptional[any]("11"),
			expectsError: false,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		},
		{
			name:         "Some Option -> Option -> option of Int should panic",
			value:        NewOptional[any](NewOptional[any](NewOptional[any](10))),
			expectsError: true,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		}, {
			name:         "None",
			value:        None[any](),
			expected:     None[any](),
			expectsError: false,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		},
		{
			name:         "Option of None",
			value:        NewOptional[any](None[any]()),
			expected:     None[any](),
			expectsError: false,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectsError == true {
				assert.Panics(t, func() {
					_ = tc.value.(Option[any]).FlatMap(tc.fn)
				})
				return
			}

			result := tc.value.(Option[any]).FlatMap(tc.fn)
			assert.Equal(t, result.Get(), tc.expected.Get())
		})
	}
}
func TestFlatMap(t *testing.T) {
	type testData[A any, B Option[any]] struct {
		name         string
		value        A
		fn           func(A any) B
		expected     Option[any]
		expectsError bool
	}

	tt := []testData[any, Option[any]]{
		{
			name:         "Some Option of Int",
			value:        NewOptional[any](10),
			expected:     NewOptional[any]("11"),
			expectsError: false,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		}, {
			name:         "Some Option of Option of Int should return option of string",
			value:        NewOptional[any](NewOptional[any](10)),
			expected:     NewOptional[any]("11"),
			expectsError: false,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		},
		{
			name:         "Some Option -> Option -> option of Int should panic",
			value:        NewOptional[any](NewOptional[any](NewOptional[any](10))),
			expectsError: true,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		}, {
			name:         "None",
			value:        None[any](),
			expected:     None[any](),
			expectsError: false,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		},
		{
			name:         "Option of None",
			value:        NewOptional[any](None[any]()),
			expected:     None[any](),
			expectsError: false,
			fn: func(a any) Option[any] {
				sum := a.(int) + 1
				return NewOptional[any](strconv.Itoa(sum))
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectsError == true {
				assert.Panics(t, func() {
					_ = FlatMap(tc.value.(Option[any]), tc.fn)
				})
				return
			}

			result := FlatMap(tc.value.(Option[any]), tc.fn)
			assert.Equal(t, result.Get(), tc.expected.Get())
		})
	}
}

func TestMap(t *testing.T) {

	type testData[A, B any] struct {
		name  string
		value A
		fn    func(value A) B
	}

	var emptySlice []string
	emptySliceAlt := []string{"some alternative slice"}

	var emptyStr = ""
	emptyVariableAlt := "some alternative string"
	tt := []testData[any, any]{
		{
			name:  "Some Number to string",
			value: 10,
			fn: func(value any) any {
				return fmt.Sprintf("%v", value)
			},
		}, {
			name:  "Some string to int",
			value: "saddam",
			fn: func(value any) any {
				str := fmt.Sprintf("%s abu ghaida", value)
				hash := sha1.New()
				hash.Write([]byte(str))
				return hash.Size()
			},
		}, {
			name:  "Some array to map",
			value: []int{1, 2, 3, 4},
			fn: func(value any) any {
				m := map[int]int{}
				for i, value := range value.([]int) {
					m[i+1] = value
				}
				return m
			},
		}, {
			name:  "Some map to array",
			value: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			fn: func(value any) any {
				v := value.(map[int]int)
				var arr []int
				for i := 0; i < 5; i++ {
					arr = append(arr, v[i])
				}

				return arr
			},
		}, {
			name: "Some struct map on name value",
			value: struct {
				name string
				age  int
			}{
				name: "saddam",
				age:  111,
			},
			fn: func(value any) any {
				v := value.(struct {
					name string
					age  int
				})

				return v.name
			},
		}, {
			name:  "Some empty string",
			value: emptyStr,
			fn: func(_ any) any {
				return emptyVariableAlt
			},
		}, {
			name:  "Some emptySlice",
			value: emptySlice,
			fn: func(_ any) any {
				return emptySliceAlt
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			option := NewOptional(tc.value)
			app := Map(option, tc.fn)
			if option.IsNone() {
				assert.True(t, app.IsNone())
				return
			}
			assert.Equal(t, option.t, app.t)
			assert.Equal(t, app.Get(), tc.fn(option.Unwrap()))
		})
	}
}
