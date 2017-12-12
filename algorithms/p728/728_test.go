package p728

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {

	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}
	assert.Equal(t, res, selfDividingNumbers(1, 22))
}
