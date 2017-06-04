package p027

import (
	"testing"
)

func test(t *testing.T, nums []int, exp []int) {
	if ans := removeElement(nums, 3); ans != len(exp) {
		t.Fatal("error answer length", ans)
	}
	for i, v := range exp {
		if v != nums[i] {
			t.Error("error answer val", nums[i], "expected:", v)
		}
	}
}

func TestExample(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	exp := []int{2, 2}
	test(t, nums, exp)
}

func TestExample1(t *testing.T) {
	nums := []int{3, 2, 2, 3, 2, 2, 3}
	exp := []int{2, 2, 2, 2}
	test(t, nums, exp)
}

func TestExample2(t *testing.T) {
	nums := []int{3, 2, 2, 3, 3, 3, 3, 2}
	exp := []int{2, 2, 2}
	test(t, nums, exp)
}

func TestExample3(t *testing.T) {
	nums := []int{3, 2, 2, 3, 2}
	exp := []int{2, 2, 2}
	test(t, nums, exp)
}

func TestExample4(t *testing.T) {
	nums := []int{3}
	exp := []int{}
	test(t, nums, exp)
}

func TestExample5(t *testing.T) {
	nums := []int{3, 3, 2}
	exp := []int{2}
	test(t, nums, exp)
}

func TestExample6(t *testing.T) {
	nums := []int{2}
	exp := []int{2}
	test(t, nums, exp)
}
