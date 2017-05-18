package p050

/**
Implement pow(x, n).
 */

func myPow(x float64, n int) float64 {
	if n == 0{
		return 1
	} else if n == 1 {
		return x
	}
	negtive := (n < 0)
	if negtive {
		n = -n
	}
	ans := myPow(x, n/2)
	ans = ans * ans
	if n % 2 != 0{
		ans = ans * x
	}
	if negtive {
		return 1 / ans
	}else {
		return ans
	}
}
