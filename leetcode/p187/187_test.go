package p187

import "testing"

func TestExample(t *testing.T) {
	ans := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	if len(ans) != 2 {
		t.FailNow()
	}
	if ans[0] != "AAAAACCCCC" {
		t.FailNow()
	}
	if ans[1] != "CCCCCAAAAA" {
		t.FailNow()
	}
}

func TestExample1(t *testing.T) {
	ans := findRepeatedDnaSequences("AAAAAAAAAAAAAAACCCCCCAAAAAGGGTTT")

	if len(ans) != 1 {
		t.FailNow()
	}
	if ans[0] != "AAAAAAAAAA" {
		t.FailNow()
	}
}
