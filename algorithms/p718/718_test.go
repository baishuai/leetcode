package p718

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	A := []int{1, 2, 3, 2, 1}
	B := []int{3, 2, 1, 4, 5}
	assert.Equal(t, 3, findLength(A, B))
}
