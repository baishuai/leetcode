package p217

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	assert.Equal(t, false, containsDuplicate([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, true, containsDuplicate([]int{1, 2, 4, 4, 5}))
}
