package p091

import "testing"

func test(t *testing.T, s string, count int) {
	if ans := numDecodings(s); ans != count {
		t.Fatal("Error answer", ans, "expected", count)
	}
}

func TestExapmle0(t *testing.T) {
	test(t, "12", 2)
}

func TestExtra0(t *testing.T) {
	test(t, "123456789", 3)
}

func TestExtra1(t *testing.T) {
	test(t, "0", 0)
}

func TestExtra2(t *testing.T) {
	test(t, "00", 0)
}

func TestExtra3(t *testing.T) {
	test(t, "27", 1)
}

func TestExtra4(t *testing.T) {
	test(t, "01", 0)
}