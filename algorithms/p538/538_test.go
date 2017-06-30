package p538

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	root := &TreeNode{Val: 0}
	exp := &TreeNode{Val: 0}
	assert.Equal(t, exp, convertBST(root))
}
