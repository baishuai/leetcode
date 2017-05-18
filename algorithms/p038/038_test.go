package p038

import (
	"testing"
)

func test(t *testing.T, n int, exp string) {
	if ans := countAndSay(n); ans != exp {
		t.Fatal("error answer", ans, "expected", exp)
	}
}

func Test1(t *testing.T) {
	test(t, 1, "1")
}

func Test2(t *testing.T) {
	test(t, 2, "11")
}

func Test3(t *testing.T) {
	test(t, 3, "21")
}

func Test4(t *testing.T) {
	test(t, 4, "1211")
}

func Test5(t *testing.T) {
	test(t, 5, "111221")
}

func Test6(t *testing.T) {
	test(t, 6, "312211")
}
