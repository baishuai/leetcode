package p029

import (
	"math/rand"
	"testing"
)

func TestRandom(t *testing.T) {
	count := 10
	var dividend, divisor int
	for count > 0 {
		count--
		dividend = rand.Int()
		divisor = rand.Int()
		if ans := divide(dividend, divisor); divisor != 0 && ans != dividend/divisor {
			t.Error("error answer", dividend, divisor, dividend/divisor)
		}

	}
}

func TestExtra0(t *testing.T) {
	if ans := divide(10, 3); ans != 10/3 {
		t.Error("error answer", 10, 3, 10/3, ans)
	}
}
