package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	iter := Empty[int]()
	assert.Equal(t, iter.Count(), 0)
	assert.Equal(t, iter.Size(), 0)
	assert.False(t, iter.HasNext())
	next := iter.Next()
	assert.Equal(t, next, 0)
}
