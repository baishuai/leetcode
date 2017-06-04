package p476

import "testing"

func Test476Example(t *testing.T) {
	if ans := findComplement(5); ans != 2 {
		t.Fatal("error answer", ans)
	}

	if findComplement(1) != 0 {
		t.Fatal("error answer")
	}
}

func Test476Cus(t *testing.T) {
	if ans := findComplement(21); ans != 10 {
		t.Fatal("error answer", ans)
	}
}
