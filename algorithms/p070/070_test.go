package p070

import "testing"

func test(t *testing.T, n, exp int) {
	if ans := climbStairs(n); ans != exp {
		t.Fatal("error answer", ans, "expected", exp)
	}
}

func Test0(t *testing.T) {
	test(t, 0, 0)
}

func Test1(t *testing.T) {
	test(t, 1, 1)
}

func Test2(t *testing.T) {
	test(t, 2, 2)
}

func Test3(t *testing.T) {
	test(t, 3, 3)
}

func Test4(t *testing.T) {
	test(t, 4, 5)
}
