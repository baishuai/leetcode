package p483

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "3", smallestGoodBase("13"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "2", smallestGoodBase("7"))
	assert.Equal(t, "2", smallestGoodBase("15"))
	assert.Equal(t, "2", smallestGoodBase("31"))
}

func Test2(t *testing.T) {
	assert.Equal(t, "59", smallestGoodBase("3541"))
}

func Test3(t *testing.T) {
	assert.Equal(t, "2", smallestGoodBase("2251799813685247"))
}
