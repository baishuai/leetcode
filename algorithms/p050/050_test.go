package p050

import (
	"math"
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 100000; i++ {
		f, n := rand.Float64(), rand.Intn(9876543210)
		if ans := myPow(f, n); ans != math.Pow(f, float64(n)) {
			t.Fatal("error answer", f, n, ans, math.Pow(f, float64(n)))
		}
	}

}

// error answer 0.09696951891448456 2775422040480279449 0 +Inf