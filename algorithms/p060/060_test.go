package p060

import "testing"

func test(t *testing.T, n, k int, exp string) {
	if ans := getPermutation(n, k); ans != exp {
		t.Fatal("error answer", ans, "expected", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, 3, 1, "123")
}

func TestExample1(t *testing.T) {
	test(t, 3, 2, "132")
}

func TestExtra0(t *testing.T) {
	test(t,5,16,"14352")
}