package p728

func isSelfDividingNumber(num int) bool {
	v := num
	for v != 0 {
		d := v % 10
		v = v / 10
		if d == 0 || num%d != 0 {
			return false
		}
	}
	return true
}

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			res = append(res, i)
		}
	}
	return res
}
