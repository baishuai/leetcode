package p606

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}}
	assert.Equal(t, "1(2(4))(3)", tree2str(root))
}

func Test1(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}
	assert.Equal(t, "1(2()(4))(3)", tree2str(root))
}

func Test2(t *testing.T) {
	var root *TreeNode
	assert.Equal(t, "", tree2str(root))
}
