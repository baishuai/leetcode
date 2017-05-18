package p207

import "testing"

func Test0(t *testing.T) {
	ans := canFinish(2, [][]int{{1, 0}})
	if !ans {
		t.Fatal(ans)
	}
}

func Test1(t *testing.T) {
	ans := canFinish(2, [][]int{{1, 0}, {0, 1}})
	if ans {
		t.Fatal(ans)
	}
}
