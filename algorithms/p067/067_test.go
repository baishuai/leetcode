package p067

import "testing"

func test(t *testing.T, a, b, s string) {
	if r := addBinary(a, b); r != s {
		t.Fatal("error answer", a, b, r, s)
	}
}

func TestExample0(t *testing.T) {
	test(t, "11", "1", "100")
}
