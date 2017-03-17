package p052

import "testing"

func test(t *testing.T, n, exp int) {
	if ans := totalNQueens(n); ans != exp {
		t.Fatal("error ", ans, "expected", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, 4, 2)
	test(t, 1, 1)
}
