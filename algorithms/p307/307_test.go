package p307

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	na := Constructor([]int{1, 3, 5})
	assert.Equal(t, 9, na.SumRange(0, 2))
	na.Update(1, 2)
	assert.Equal(t, 8, na.SumRange(0, 2))
}
