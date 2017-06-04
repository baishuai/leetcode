package p152

import "testing"

func TestExample(t *testing.T) {
	if maxProduct([]int{2, 3, -2, 4}) != 6 {
		t.Fatal("error answer")
	}
}

func TestExtra0(t *testing.T) {
	if maxProduct([]int{-2}) != -2 {
		t.Fatal("error answer")
	}
}

func TestExtra1(t *testing.T) {
	if maxProduct([]int{-3, -4}) != 12 {
		t.Fatal("error answer")
	}
}

func TestExtra2(t *testing.T) {
	if ans := maxProduct([]int{-3, 2, -4}); ans != 24 {
		t.Fatal("error answer", ans)
	}
}

func TestExtra3(t *testing.T) {
	if ans := maxProduct([]int{-3, 2, -4, 0, 4, 8}); ans != 32 {
		t.Fatal("error answer", ans)
	}
}
