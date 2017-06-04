package p040

import "testing"

func test(t *testing.T, cand []int, target int, exp [][]int) {
	ans := combinationSum2(cand, target)
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
	cand := []int{10, 1, 2, 7, 6, 1, 5}
	exp := [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	}
	test(t, cand, 8, exp)
}

func TestExtra0(t *testing.T) {
}
