package p042

import "testing"

func test(t *testing.T, height []int, exp int) {
	if ans := trap(height); ans != exp {
		t.Fatal("error answer", ans)
	}
}

func TestExample1(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	test(t, height, 6)
}

func TestExtra1(t *testing.T) {
	height := []int{0, 1, 2, 3, 4, 5}
	test(t, height, 0)
}

func TestExtra2(t *testing.T) {
	height := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1}
	test(t, height, 0)
}

func TestExtra3(t *testing.T) {
	height := []int{5, 4, 3, 2, 1}
	test(t, height, 0)
}

func TestExtra4(t *testing.T) {
	height := []int{5, 4, 3, 2, 4}
	test(t, height, 3)
}

func TestExtra5(t *testing.T) {
	height := []int{3, 0, 5, 3, 3, 6, 5, 0, 3, 6}
	test(t, height, 17)
}

func TestExtra6(t *testing.T) {
	height := []int{5, 2, 1, 2, 1, 5}
	test(t, height, 14)
}
