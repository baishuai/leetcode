package p007

import (
	"testing"
)

func TestInt(t *testing.T) {
	result := reverse(123)
	if result != 321 {
		t.Fatal("error anwser", result)
	}
}


func TestZero(t *testing.T) {
	result := reverse(0)
	if result != 0 {
		t.Fatal("error anwser")
	}
}



func TestNeg(t *testing.T) {
	result := reverse(-123)
	if result != -321 {
		t.Fatal("error anwser", result)
	}
}


func TestOverflow(t *testing.T) {
	result := reverse(9223372036854775599)
	if result != 0 {
		t.Fatal("error anwser")
	}
}

func TestNotOverflow(t *testing.T) {
	result := reverse(1534236469)
	if result != 0 {
		t.Fatal("error anwser", result)
	}
}

func TestNegConf(t *testing.T) {
	result := reverse(-2147483648)
	if result != 0 {
		t.Fatal("error anwser", result)
	}
}
