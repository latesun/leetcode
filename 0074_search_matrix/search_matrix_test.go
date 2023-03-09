package searchmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	cases := []struct {
		input    [][]int
		target   int
		expected bool
	}{
		{
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expected, searchMatrix(cases[i].input, cases[i].target))
	}
}
