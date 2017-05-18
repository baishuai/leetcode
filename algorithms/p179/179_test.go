package p179

import "testing"

func TestExample(t *testing.T) {
	ans := largestNumber([]int{3, 30, 34, 5, 9})
	if ans != "9534330" {
		t.Fatal(ans, "9534330")
	}
}

func TestExample1(t *testing.T) {
	ans := largestNumber([]int{1, 2, 3, 4, 45, 65, 34, 42, 67, 8, 76, 8, 6})
	if ans != "8876676654544234321" {
		t.Fatal(ans, "8876676654544234321")
	}
}

func TestExtra(t *testing.T) {
	ans := largestNumber([]int{121, 12})
	if ans != "12121" {
		t.Fatal(ans, "12121")
	}
}

func TestExtra2(t *testing.T) {
	ans := largestNumber([]int{0, 0, 0})
	if ans != "0" {
		t.Fatal(ans, "0")
	}
}
