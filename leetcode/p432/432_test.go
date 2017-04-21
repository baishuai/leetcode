package p432

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	allone := Constructor()

	allone.Inc("a")

	allone.Inc("b")
	allone.Inc("a")
	assert.Equal(t, "a", allone.GetMaxKey())
	assert.Equal(t, "b", allone.GetMinKey())
	allone.Dec("a")
	allone.Dec("a")
	assert.Equal(t, "b", allone.GetMinKey())
	assert.Equal(t, "b", allone.GetMaxKey())

}
