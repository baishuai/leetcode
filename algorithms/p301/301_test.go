package p301

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	//"()())()" -> ["()()()", "(())()"]
	ans := removeInvalidParentheses("()())()")
	assert.Equal(t, []string{"(())()", "()()()"}, ans)
}

func Test1(t *testing.T) {
	assert.Equal(t, []string{""}, removeInvalidParentheses("("))
}

func Test2(t *testing.T) {
	assert.Equal(t, []string{""}, removeInvalidParentheses(""))
}

func Test3(t *testing.T) {
	assert.Equal(t, []string{"f"}, removeInvalidParentheses(")(f"))
}
