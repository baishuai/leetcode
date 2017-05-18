package p122

import "testing"

func test(t *testing.T, prices []int, exp int) {
	if ans := maxProfit(prices); ans != exp {
		t.Fatal("error answer")
	}
}

func TestEXample0(t *testing.T) {
	test(t, []int{7, 1, 5, 3, 6, 4}, 7)
}

func TestExample1(t *testing.T) {
	test(t, []int{7, 6, 4, 3, 1}, 0)
}
