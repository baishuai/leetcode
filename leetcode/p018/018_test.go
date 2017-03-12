package p018

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	nums := []int{1, 0, -1, 0, 0, -2, 2}
	ans := fourSum(nums, 0)
	for _, v := range ans {
		fmt.Println(v)
	}
}

func TestExtra0(t *testing.T) {
	nums := []int{0,0,0,0}
	ans := fourSum(nums, 0)
	for _, v := range ans {
		fmt.Println(v)
	}
}
