package p217

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, false, containsDuplicate([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, true, containsDuplicate([]int{1, 2, 4, 4, 5}))
}
