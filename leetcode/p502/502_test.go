package p502

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	p := []int{1, 2, 3}
	c := []int{0, 1, 1}
	assert.Equal(t, 4, findMaximizedCapital(2, 0, p, c))
}

func Test1(t *testing.T) {
	p := []int{1, 2, 3}
	c := []int{0, 1, 2}
	assert.Equal(t, 6, findMaximizedCapital(10, 0, p, c))
}
