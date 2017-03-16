package p135

import "testing"

func test(t *testing.T, ratings []int, exp int) {
	if ans := candy(ratings); ans != exp {
		t.Fatal("error", ans, "expected", exp)
	}

	if ans := candy_another(ratings); ans != exp {
		t.Fatal("error", ans, "expected", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, []int{0}, 1)
}

func TestExample1(t *testing.T) {
	test(t, []int{0, 0}, 2)
}

func TestExample2(t *testing.T) {
	test(t, []int{4, 3, 2, 1, 0}, 15)
}

func TestExample3(t *testing.T) {
	test(t, []int{1, 2, 4, 3, 5, 2, 0, 1}, 15)
}

func TestExample4(t *testing.T) {
	test(t, []int{1, 2, 2}, 4)
}

func TestExample5(t *testing.T) {
	test(t, []int{1, 2, 1}, 4)
}

func TestExtra0(t *testing.T) {
	test(t, []int{5, 1, 1, 1, 10, 2, 1, 1, 1, 3}, 15)
}
