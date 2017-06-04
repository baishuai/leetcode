package p010

import "testing"

func TestExample(t *testing.T) {

	if isMatch("aa", "a") {
		t.Fatal("error answer")
	}

	if !isMatch("aa", "aa") {
		t.Fatal("error answer")
	}

	if isMatch("aaa", "aa") {
		t.Fatal("error answer")
	}

	if !isMatch("aa", "a*") {
		t.Fatal("error answer")
	}

	if !isMatch("aa", ".*") {
		t.Fatal("error answer")
	}

	if !isMatch("ab", ".*") {
		t.Fatal("error answer")
	}

}

func TestAab(t *testing.T) {
	if !isMatch("aab", "c*a*b") {
		t.Fatal("error answer")
	}
}

func TestA_a(t *testing.T) {
	if !isMatch("aaa", "a*a") {
		t.Fatal("error answer")
	}
}
