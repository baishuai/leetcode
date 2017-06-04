package p514

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	ans := findRotateSteps("godding", "gd")
	assert.Equal(t, 4, ans)
}

func Test1(t *testing.T) {
	assert.Equal(t, 10, findRotateSteps("edcba", "abcde"))
}
