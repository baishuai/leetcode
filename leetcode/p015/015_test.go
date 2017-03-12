package p015

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	for _, v := range ans {
		fmt.Println(v)
	}
}
