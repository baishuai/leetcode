package p009

import "testing"

func Test9Zero(t *testing.T)  {
	if !isPalindrome(0) {
		t.Fatal("error anwser")
	}
}

func Test9MZero(t *testing.T)  {
	if !isPalindrome(11) {
		t.Fatal("error anwser")
	}
}

func Test31Zero(t *testing.T)  {
	if !isPalindrome(111) {
		t.Fatal("error anwser")
	}
}

func Test9M1Zero(t *testing.T)  {
	if !isPalindrome(121) {
		t.Fatal("error anwser")
	}
}

func Test9lMZero(t *testing.T)  {
	if isPalindrome(1011) {
		t.Fatal("error anwser")
	}
}

func Test9rMZero(t *testing.T)  {
	if isPalindrome(1101) {
		t.Fatal("error anwser")
	}
}