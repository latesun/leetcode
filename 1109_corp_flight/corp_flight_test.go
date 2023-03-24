package corpflight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	bookings [][]int
	n        int
	expected []int
}{
	{
		[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
		5,
		[]int{10, 55, 45, 25, 25},
	},
}

func TestCorpFlightBookings(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expected, corpFlightBookings(tt.bookings, tt.n))
	}
}
