package p102

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {
	assert.Equal(t, [][]int{}, levelOrder(nil))
}

func Test1(t *testing.T) {
	assert.Equal(t, [][]int{{1}, {2, 3}}, levelOrder(&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}))
}
