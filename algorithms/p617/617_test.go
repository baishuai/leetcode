package p617

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {

	assert.Nil(t, mergeTrees(nil, nil))
	assert.Equal(t, &TreeNode{Val: 1}, mergeTrees(nil, &TreeNode{Val: 1}))
	assert.Equal(t, &TreeNode{Val: 1}, mergeTrees(&TreeNode{Val: 1}, nil))
}

func Test1(t *testing.T) {

	assert.Equal(t, &TreeNode{1, nil, &TreeNode{Val: 1}}, mergeTrees(nil, &TreeNode{1, nil, &TreeNode{Val: 1}}))
	assert.Equal(t, &TreeNode{1, nil, &TreeNode{Val: 1}}, mergeTrees(&TreeNode{1, nil, &TreeNode{Val: 1}}, nil))

	assert.Equal(t, &TreeNode{1, &TreeNode{Val: 1}, nil}, mergeTrees(nil, &TreeNode{1, &TreeNode{Val: 1}, nil}))
	assert.Equal(t, &TreeNode{1, &TreeNode{Val: 1}, nil}, mergeTrees(&TreeNode{1, &TreeNode{Val: 1}, nil}, nil))
}
