package p172

import "testing"

func test(t *testing.T, n int, cnt int) {
	if ans := trailingZeroes(n); ans != cnt {
		t.Fatal(ans, "expected", cnt)
	}
}

func Test1(t *testing.T) {
	test(t, 23, 4)
}

func Test0(t *testing.T) {
	test(t, 0, 0)
}

func Test100(t *testing.T) {
	test(t, 100, 24)
}
