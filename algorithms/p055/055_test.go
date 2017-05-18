package p055

import "testing"

func test(t *testing.T, nums []int, exp bool)  {
	if canJump(nums) != exp{
		t.Fatal("error answer")
	}
}

func TestExample0(t *testing.T) {
	test(t, []int{2,3,1,1,4},true)
}


func TestExample1(t *testing.T) {
	test(t, []int{3,2,1,0,4},false)
}