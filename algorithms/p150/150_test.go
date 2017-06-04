package p150

import "testing"

func test(t *testing.T, tokens []string, exp int) {
	if ans := evalRPN(tokens); ans != exp {
		t.Fatal("error", ans, "Exp", exp)
	}
}

func TestEXample0(t *testing.T) {
	test(t, []string{"2", "1", "+", "3", "*"}, 9)
}

func TestEXample1(t *testing.T) {
	test(t, []string{"4", "13", "5", "/", "+"}, 6)
}
