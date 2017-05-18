package p125

import "testing"

func TestExample0(t *testing.T) {
	if !isPalindrome("A man, a plan, a canal: Panama") {
		t.Fail()
	}
}

func TestExample1(t *testing.T) {
	if isPalindrome("race a car") {
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	if !isPalindrome("r") {
		t.Fail()
	}
}

func TestExample3(t *testing.T) {
	if !isPalindrome("") {
		t.Fail()
	}
}

func TestExample4(t *testing.T) {
	if isPalindrome("01W") {
		t.Fail()
	}
}

func TestExample5(t *testing.T) {
	if isPalindrome("ab2a") {
		t.Fail()
	}
}
