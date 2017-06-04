package p069

import "testing"

func test(t *testing.T, x, s int) {
	if mySqrt(x) != s {
		t.Fatal("error ans", x, s)
	}
}

func TestExample(t *testing.T) {
	test(t, 10, 3)
}

func TestExtra0(t *testing.T) {
	test(t, 0, 0)
}
