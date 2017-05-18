package p137

import "testing"

func TestExample(t *testing.T) {
	ans := singleNumber([]int{1, 1, 2, 2, 3, 3, 4, 4, 6, 6, 7, 7, 7, 6, 5, 4, 3, 2, 1})
	if ans != 5 {
		t.Fatal(ans, 5)
	}
}
