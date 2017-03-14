package p119

import "testing"

func TestExample(t *testing.T) {
	ans := getRow(3)
	exp := []int{1, 3, 3, 1}
	if len(ans) != len(exp) {
		t.Fatal("error answer length")
	}
	for i, v := range exp {
		if ans[i] != v {
			t.Error("error value", i, ans[i], v)
		}
	}
}
