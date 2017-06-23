package p625

import (
	"math"
)

func smallestFactorization(a int) int {

	if a < 10 {
		return a
	}
	facs := make([]int, 0)

	for i := 9; i > 1; i-- {
		for a%i == 0 {
			a = a / i
			facs = append(facs, i)
		}
	}

	if a > 10 {
		return 0
	}

	ans := int64(0)

	for i := len(facs) - 1; i >= 0; i-- {
		ans = ans*10 + int64(facs[i])
	}
	if ans > math.MaxInt32 {
		ans = 0
	}
	return int(ans)
}
