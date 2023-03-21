package removeduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		input             []int
		nonDuplicateCount int
	}{
		{
			[]int{1, 1, 2},
			2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
	}

	for i := range cases {
		assert.Equal(t, cases[i].nonDuplicateCount, removeDuplicates(cases[i].input))
	}
}
