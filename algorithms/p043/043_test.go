package p043

import "testing"

func test(t *testing.T, num1 string, num2 string, exp string) {
	if ans := multiply(num1, num2); ans != exp {
		t.Fatal("error answer", ans, "expected", exp)
	}
}

func TestExample(t *testing.T) {
	test(t, "0", "0", "0")
}

func TestExtra0(t *testing.T) {
	test(t, "00", "00000", "0")
}

func TestExtra1(t *testing.T) {
	test(t, "123", "11", "1353")
}

func TestExtra2(t *testing.T) {
	test(t, "9", "9", "81")
}