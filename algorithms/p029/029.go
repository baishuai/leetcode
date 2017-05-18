package p029

import "math"

/**
Divide two integers without using multiplication,
 division and mod operator.

If it is overflow, return MAX_INT.
 */

func absi64(x int64) int64 {
	if x < 0 {
		x = -x
	} else if x == 0 {
		x = 0
	}
	return x
}

func divide(dividend int, divisor int) int {

	flag := (dividend > 0 ) == (divisor > 0)
	ans := int64(0)

	dividend64 := absi64(int64(dividend))
	divisor64 := absi64(int64(divisor))

	if divisor64 == 0 {
		return math.MaxInt32
	}
	if dividend64 == 0 || (dividend64 < divisor64) {
		return 0
	}

	for dividend64 >= divisor64 {
		temp, mul := divisor64, int64(1)
		for dividend64 >= (temp << 1) {
			temp <<= 1
			mul <<= 1
		}
		dividend64 -= temp
		ans += mul
	}

	if ans > math.MaxInt32 {
		if flag {
			ans = math.MaxInt32
		} else {
			ans = math.MinInt32
		}
	} else {
		if !flag {
			ans = - ans
		}
	}
	return int(ans)
}
