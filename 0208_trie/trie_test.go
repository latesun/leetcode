package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	obj := Constructor()
	obj.Insert("hi")
	obj.Insert("hello")

	assert.Equal(t, true, obj.Search("hi"))
	assert.Equal(t, false, obj.Search("ih"))
	assert.Equal(t, true, obj.StartsWith("he"))
}
