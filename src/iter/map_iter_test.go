package iter

import (
	"github.com/stretchr/testify/assert"
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
