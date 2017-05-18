package p461

import "testing"

func Test461Example(t *testing.T) {
	ans := hammingDistance(1, 4)
	if ans != 2 {
		t.Fatal("error answer", ans)
	}
}
