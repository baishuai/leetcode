package p377

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 7, combinationSum4([]int{1, 2, 3}, 4))
}
