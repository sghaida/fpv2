package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	iter := Empty[int]()
	assert.Equal(t, iter.Count(), 0)
	assert.Equal(t, iter.Size(), 0)
	next, ok := iter.Next()
	assert.False(t, ok)
	assert.Equal(t, next, 0)
}
