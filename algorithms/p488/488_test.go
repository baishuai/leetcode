package p488

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, 2, findMinStep("WWRRBBWW", "WRBRW"))
}

func Test1(t *testing.T) {
	assert.Equal(t, -1, findMinStep("WRRBBW", "RB"))
}

func Test2(t *testing.T) {
	assert.Equal(t, 2, findMinStep("G", "GGGG"))
}

func Test3(t *testing.T) {
	assert.Equal(t, 3, findMinStep("RBYYBBRRB", "YRBGB"))
}
