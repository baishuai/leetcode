package p058

import "testing"

func TestExample(t *testing.T) {
	if lengthOfLastWord(" nnnn  ") != 4 {
		t.Fatal("error answer")
	}
}
