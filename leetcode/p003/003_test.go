package p003

import "testing"

func TestBbb(t *testing.T) {
	var max int
	max = lengthOfLongestSubstring("bbb")
	if max != 1 {
		t.Fatal("err result, should be 1, but", max)
	}


	max = lengthOfLongestSubstring("qwertyhhyjgb")
	if max != 7 {
		t.Fatal("err result, should be 7, but", max)
	}
}
