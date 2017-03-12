package p072

import "testing"

func test(t *testing.T, w1, w2 string, exp int) {
	if ans := minDistance(w1, w2); ans != exp {
		t.Fatal("error answer", ans, "expected", exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, "a", "a", 0)
}

func TestExample1(t *testing.T) {
	test(t, "", "", 0)
}

func TestExample2(t *testing.T) {
	test(t, "a", "b", 1)
}

func TestExample3(t *testing.T) {
	test(t, "ab", "b", 1)
}

func TestExample4(t *testing.T) {
	test(t, "aabbcc", "b", 5)
}

func TestExample5(t *testing.T) {
	test(t, "aabbcc", "abc", 3)
}

func TestExample6(t *testing.T) {
	test(t, "", "abc", 3)
}

func TestExample7(t *testing.T) {
	test(t, "abc", "", 3)
}