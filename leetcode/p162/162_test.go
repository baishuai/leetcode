package p162

import "testing"

func TestExample(t *testing.T) {
	ans := findPeakElement([]int{1, 2, 3, 1})
	if ans != 2 {
		t.Fatal(ans, 2)
	}
}

func TestExample1(t *testing.T) {
	ans := findPeakElement([]int{1, 2, 3})
	if ans != 2 {
		t.Fatal(ans, 2)
	}
}

func TestExample2(t *testing.T) {
	ans := findPeakElement([]int{2, 3, 4, 1})
	if ans != 2 {
		t.Fatal(ans, 2)
	}
}
