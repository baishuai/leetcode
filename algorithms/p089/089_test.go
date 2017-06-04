package p089

import "testing"

func test(t *testing.T, n int, exp []int) {
	ans := grayCode(n)
	if len(ans) != len(exp) {
		t.Fatal("error answer length")
	}
	for i, v := range exp {
		if ans[i] != v {
			t.Error("error answer value", i, ans[i], v)
		}
	}
}

func TestExample0(t *testing.T) {
	test(t, 0, []int{0})
}

func TestExample1(t *testing.T) {
	test(t, 1, []int{0, 1})
}

func TestExample2(t *testing.T) {
	test(t, 2, []int{0, 1, 3, 2})
}

func TestExample3(t *testing.T) {
	test(t, 3, []int{0, 1, 3, 2, 6, 7, 5, 4})
}
