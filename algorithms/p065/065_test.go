package p065

import "testing"

func test(t *testing.T, s string, exp bool) {
	if ans := isNumber(s); ans != exp {
		t.Fatal("error", s, ans, "expected", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, "0", true)
}

func TestExample1(t *testing.T) {
	test(t, " 0.1", true)
}

func TestExample2(t *testing.T) {
	test(t, "abs", false)
}

func TestExample3(t *testing.T) {
	test(t, "1 a", false)
}

func TestExample4(t *testing.T) {
	test(t, "2e10", true)
}

func TestExtra0(t *testing.T) {
	test(t, ".", false)
}

func TestExtra1(t *testing.T) {
	test(t, ".2", true)
}

func TestExtra2(t *testing.T) {
	test(t, "e", false)
}

func TestExtra3(t *testing.T) {
	test(t, "2e", false)
}

func TestExtra4(t *testing.T) {
	test(t, "2.", true)
}

func TestExtra5(t *testing.T) {
	test(t, " ", false)
}

func TestExtra6(t *testing.T) {
	test(t, "-2", true)
}

func TestExtra7(t *testing.T) {
	test(t, "-2.3e45", true)
}

func TestExtra8(t *testing.T) {
	test(t, "e45", false)
}

func TestExtra9(t *testing.T) {
	test(t, ".3e45", true)
}

func TestExtra10(t *testing.T) {
	test(t, "0e", false)
}
