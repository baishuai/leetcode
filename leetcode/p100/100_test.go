package p100

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, true, isSameTree(nil, nil))
}

func Test2(t *testing.T) {
	assert.Equal(t, true, isSameTree(&TreeNode{}, &TreeNode{}))
}

func Test3(t *testing.T) {
	assert.Equal(t, false, isSameTree(nil, &TreeNode{}))
}
