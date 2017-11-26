package p668

func findKthNumber(m int, n int, k int) int {
	lo := 1
	hi := m * n
	for lo < hi {
		mi := lo + (hi-lo)/2
		if !enough(mi, m, n, k) {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo
}

func enough(x, m, n, k int) bool {
	count := 0
	for i := 1; i <= m; i++ {
		count += min(x/i, n)
	}
	return count >= k
}

func min(a, b int) int {
	if a < b {
		b = a
	}
	return b
}
