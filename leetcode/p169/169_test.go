package p169

import "testing"

func TestExample0(t *testing.T) {
	if ans := majorityElement([]int{1}); ans != 1 {
		t.FailNow()
	}
}

func TestExample1(t *testing.T) {
	if ans := majorityElement([]int{1, 2, 1}); ans != 1 {
		t.FailNow()
	}
}
