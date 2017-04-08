package p213

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func test(t *testing.T, nums []int, exp int) {
	ans := rob(nums)
	assert.Equal(t, exp, ans)
}

func TestExample0(t *testing.T) {
	test(t, []int{}, 0)
	test(t, nil, 0)
}

func TestExample1(t *testing.T) {
	test(t, []int{1}, 1)
	test(t, []int{1, 2}, 2)
	test(t, []int{1, 2, 0}, 2)
}
