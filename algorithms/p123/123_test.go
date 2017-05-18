package p123

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 6, maxProfit([]int{1, 2, 3, 4, 5, 6, 7}))
}
