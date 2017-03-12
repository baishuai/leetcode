package p039

import "testing"

func test(t *testing.T, cand []int, target int, exp [][]int) {
	ans := combinationSum(cand, target)
	if len(ans) != len(exp) {
		t.Fatal("error answer length", len(ans))
	}
	for i, v := range ans {
		expv := exp[i]
		if len(v) != len(expv) {
			t.Fatal("error answer length", len(ans), v, expv)
		}
		for j := 0; j < len(v); j++ {
			if v[j] != expv[j] {
				t.Fatal("error answer value", v[j], expv[j])
			}
		}
	}

}

func TestExample(t *testing.T) {
	cand := []int{2, 3, 6, 7}
	exp := [][]int{
		{2, 2, 3},
		{7},
	}
	test(t, cand, 7, exp)
}

func TestExtra0(t *testing.T) {
	cand := []int{1, 2}
	exp := [][]int{{1, 1, 1, 1}, {1, 1, 2}, {2, 2}}
	test(t, cand, 4, exp)
}
