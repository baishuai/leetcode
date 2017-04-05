package p204

import "testing"

func test(t *testing.T, n int, exp int) {
	if ans := countPrimes(n); ans != exp {
		t.Fatal(ans, "exp", exp)
	}
}

func Test0(t *testing.T) {
	test(t, 1, 0)
	test(t, 2, 0)
	test(t, 3, 1)
}

func Test10(t *testing.T) {
	test(t, 10, 4)
}
