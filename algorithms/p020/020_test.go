package p020

import "testing"

func Test20Example(t *testing.T) {
	if ! isValid("()") {
		t.Fatal("error ans")
	}

	if ! isValid("()[]{}") {
		t.Fatal("error ans")
	}
}

func Test20Wrong(t *testing.T) {
	if  isValid("()[") {
		t.Fatal("error ans")
	}

	if isValid("()[){}") {
		t.Fatal("error ans")
	}

	if isValid("({") {
		t.Fatal("error ans")
	}
}
