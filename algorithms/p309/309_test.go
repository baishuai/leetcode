package p309

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 3, maxProfit([]int{1, 2, 3, 0, 2}))
}

func Test1(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{2}))
}
