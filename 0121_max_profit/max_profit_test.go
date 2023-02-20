package maxprofit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	cases := map[int][]int{
		5: {7, 1, 5, 3, 6, 4},
		0: {7, 6, 4, 3, 1},
	}

	for expected, input := range cases {
		assert.Equal(t, expected, maxProfit(input))
	}
}
