package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := Constructor()
	assert.Equal(t, true, q.Empty())

	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, false, q.Empty())
	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Peek())
}
