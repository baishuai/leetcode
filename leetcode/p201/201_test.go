package p201

import "testing"

func test(t *testing.T, m, n int, exp int) {
	if ans := rangeBitwiseAnd(m, n); ans != exp {
		t.Fatal(ans, exp)
	}
}

func TestExample(t *testing.T) {
	test(t, 2, 3, 2)
	test(t, 3, 3, 3)
	test(t, 3, 4, 0)
	test(t, 20000, 2147483647, 0)
}

func TestExample1(t *testing.T) {
	test(t, 1, 1, 1)
}
