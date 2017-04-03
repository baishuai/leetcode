package p189

import "testing"

func TestName(t *testing.T) {
	ans := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(ans, 3)

	exp := []int{5, 6, 7, 1, 2, 3, 4}
	for i, v := range exp {
		if ans[i] != v {
			t.FailNow()
		}
	}
}
