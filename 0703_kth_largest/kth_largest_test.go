package kthlargest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthlargest(t *testing.T) {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, kthLargest.Add(3))  // [0, 4, 5, 8]
	assert.Equal(t, 5, kthLargest.Add(5))  // [0, 5, 5, 8]
	assert.Equal(t, 5, kthLargest.Add(10)) // [0, 5, 10, 8]
	assert.Equal(t, 8, kthLargest.Add(9))  // [0, 8, 10, 9]
	assert.Equal(t, 8, kthLargest.Add(4))  // [0, 8, 10, 9]
}
