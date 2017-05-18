package p168

import "testing"

func test(t *testing.T, n int, exp string) {
	if ans := convertToTitle(n); ans != exp {
		t.Fatal("error", ans, "exp", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, 28, "AB")
}

func TestExample1(t *testing.T) {
	test(t, 26, "Z")
	test(t, 1, "A")
}
