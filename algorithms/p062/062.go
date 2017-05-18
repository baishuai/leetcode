package p062

func uniquePaths(m int, n int) int {
	//组合数学 格路问题
	s := m + n - 2
	k := m - 1
	if n-1 < k {
		k = n - 1
	}
	prods, prodk := 1, 1
	for k > 0 {
		prodk *= k
		prods *= s
		s--
		k--
	}
	return prods / prodk
}
