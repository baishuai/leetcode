package p211

import "testing"
import "github.com/stretchr/testify/assert"

func Test0(t *testing.T) {
	dict := Constructor()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")
	assert.Equal(t, false, dict.Search("pad"))
	assert.Equal(t, true, dict.Search("bad"))
	assert.Equal(t, true, dict.Search(".ad"))
	assert.Equal(t, true, dict.Search("b.."))
}
