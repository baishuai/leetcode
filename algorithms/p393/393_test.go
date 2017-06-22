package p393

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, validUtf8([]int{197, 130, 1}))
}

func Test1(t *testing.T) {
	assert.Equal(t, false, validUtf8([]int{235, 140, 4}))
	assert.Equal(t, false, validUtf8([]int{0x88, 0x88, 4}))
}

func Test2(t *testing.T) {
	assert.Equal(t, false, validUtf8([]int{237}))
}
