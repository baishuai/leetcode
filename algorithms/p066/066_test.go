package p066

import "testing"

func test(t *testing.T, digits []int, exp []int) {
	ans := plusOne(digits)
	if len(ans) != len(exp) {
		t.Fatal("error answer length", len(ans))
	}
	for i, v := range ans {
		if v != exp[i] {
			t.Error("error answer", i, v, exp[i])
		}
	}
}

func TestExtra0(t *testing.T) {
	test(t, []int{0}, []int{1})
}

func TestExtra1(t *testing.T) {
	test(t, []int{9}, []int{1, 0})
}
