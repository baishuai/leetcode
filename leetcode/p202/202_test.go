package p202

import "testing"

func test(t *testing.T, n int, exp bool) {
	if isHappy(n) != exp {
		t.Fatal(!exp, "exp", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, 19, true)
}

func TestExample2(t *testing.T) {
	test(t, 2, false)
}
