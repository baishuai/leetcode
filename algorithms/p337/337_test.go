package p337

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 2, rob(&TreeNode{Val: 2}))
	assert.Equal(t, 2, rob(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))
}
