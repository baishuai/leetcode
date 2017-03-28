package p164

import "testing"

func TestExample(t *testing.T) {
	ans := maximumGap([]int{2, 5, 9, 1, 4, 6})
	if ans != 3 {
		t.Fatal("error", ans, "Exp", 3)
	}
}
