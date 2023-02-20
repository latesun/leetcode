package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	input := []int{3, 4, 5, 1, 2, 7, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	BubbleSort(input)
	assert.Equal(t, expected, input)
}

func TestInsertSort(t *testing.T) {
	input := []int{3, 4, 5, 1, 2, 7, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	InsertSort(input)
	assert.Equal(t, expected, input)
}
