package p044

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

	if ! isMatch("aa", "a*") {
		t.Fatal("error answer")
	}

	if ! isMatch("aa", "?*") {
		t.Fatal("error answer")
	}

	if ! isMatch("ab", "?*") {
		t.Fatal("error answer")
	}

	if !isMatch("a", "?*") {
		t.Fatal("error answer")
	}

	if ! isMatch("ab", "?*") {
		t.Fatal("error answer")
	}

	if ! isMatch("aad", "*") {
		t.Fatal("error answer")
	}
}

func TestAab(t *testing.T) {
	if isMatch("aab", "c*a*b") {
		t.Fatal("error answer")
	}
}

func TestA_a(t *testing.T) {
	if ! isMatch("aaa", "a*a") {
		t.Fatal("error answer")
	}
}

func TestExtra0(t *testing.T) {

	if !isMatch("abefcdgiescdfimde", "ab*cd?i*de") {
		t.Fatal("error answer")
	}
}

func TestExtra1(t *testing.T) {

	if isMatch("acefcdgiescdfimde", "ab*cd?i*de") {
		t.Fatal("error answer")
	}
}


func TestExtra2(t *testing.T) {

	if isMatch("", "a") {
		t.Fatal("error answer")
	}
}
