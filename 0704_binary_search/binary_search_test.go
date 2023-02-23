package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		input    []int
		target   int
		expected int
	}{
		{
			input:    []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			input:    []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
	}

	for i := range cases {
		assert.Equal(t, cases[i].expected, search(cases[i].input, cases[i].target))
	}
}
