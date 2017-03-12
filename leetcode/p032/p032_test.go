package p032

import (
	"testing"
)

func test(t *testing.T, s string, exp int) {
	if ans := longestValidParentheses(s); ans != exp {
		t.Fatal("error answer", ans, "expected ", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, "(()", 2)
}

func TestExample1(t *testing.T) {
	test(t, ")()())", 4)
}

func TestExtra0(t *testing.T) {
	test(t, "()", 2)
}

func TestExtra1(t *testing.T) {
	test(t, ")", 0)
}

func TestExtra2(t *testing.T) {
	test(t, "(", 0)
}

func TestExtra3(t *testing.T) {
	test(t, "()(()", 2)
}
