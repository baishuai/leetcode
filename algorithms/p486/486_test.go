package p486

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, false, PredictTheWinner([]int{1, 5, 2}))
	assert.Equal(t, true, PredictTheWinner([]int{1, 5, 23, 7}))
}
