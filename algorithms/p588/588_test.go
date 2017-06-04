package p588

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, []string{"", "dd", "ss", "ss"}, strings.Split("/dd/ss/ss", "/"))
}

func Test1(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, []string{}, obj.Ls("/"))
	obj.Mkdir("/a/b/c")
	obj.AddContentToFile("/a/b/c/d", "hello")
	assert.Equal(t, "hello", obj.ReadContentFromFile("/a/b/c/d"))
}
