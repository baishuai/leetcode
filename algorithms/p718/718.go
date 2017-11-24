package p718

func findLength(A []int, B []int) int {
	dp := make([]int, len(A))
	dpPre := make([]int, len(A))
	m := 0
	for i := 0; i < len(B); i++ {
		dpPre, dp = dp, dpPre
		for j := 0; j < len(A); j++ {
			if A[j] == B[i] {
				if i == 0 || j == 0 {
					dp[j] = 1
				} else {
					dp[j] = dpPre[j-1] + 1
				}
				if dp[j] > m {
					m = dp[j]
				}
			} else {
				dp[j] = 0
			}
		}
	}
	return m
}
