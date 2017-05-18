package p096

import "testing"

func test(t *testing.T, n, exp int) {
	if ans := numTrees(n); ans != exp {
		t.Fatal("Error answer", ans, exp)
	}
}

func TestExample0(t *testing.T) {
	test(t, 0, 1)
}

func TestExample1(t *testing.T) {
	test(t, 1, 1)
}

func TestExample2(t *testing.T) {
	test(t, 2, 2)
}

func TestExample3(t *testing.T) {
	test(t, 3, 5)
}

func TestExample4(t *testing.T) {
	test(t, 4, 14)
}
