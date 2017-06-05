package p474

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 4, findMaxForm([]string{"10", "001", "111001", "1", "0"}, 5, 3))
}

func Test1(t *testing.T) {
	assert.Equal(t, 2, findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

func Test2(t *testing.T) {
	assert.Equal(t, 0, findMaxForm([]string{"00", "000"}, 1, 10))
}
