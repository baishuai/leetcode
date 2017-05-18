package p084

import "testing"

func test(t *testing.T, heights []int, exp int) {
	if ans := largestRectangleArea(heights); ans != exp {
		t.Fatal("error", ans, "expected", exp)
	}

	if ans := largestRectangleArea_Slow(heights); ans != exp {
		t.Fatal("error", ans, "expected", exp)
	}
}

func TestExample(t *testing.T) {
	test(t, []int{2, 1, 5, 6, 2, 3}, 10)
}

func TestExtra0(t *testing.T) {
	test(t, []int{2, 1, 2}, 3)
}

func TestExtra1(t *testing.T) {
	test(t, []int{1, 1, 1, 1, 1}, 5)
}

func TestExtra2(t *testing.T) {
	test(t, []int{0}, 0)
}
