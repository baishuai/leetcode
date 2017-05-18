package p018

import (
	"testing"
)

func TestExample(t *testing.T) {
	nums := []int{1, 0, -1, 0, 0, -2, 2}
	fourSum(nums, 0)

}

func TestExtra0(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	fourSum(nums, 0)
}
