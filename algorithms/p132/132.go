package p132

/**
Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.

For example, given s = "aab",
Return 1 since the palindrome partitioning ["aa","b"] could be produced using 1 cut.
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCut(s string) int {

	//O(N^2)time O(N^2)space
	length := len(s)
	if length <= 1 {
		return 0
	}

	isPa := make([][]bool, length)
	for i := 0; i < length; i++ {
		isPa[i] = make([]bool, length)
	}
	cuts := make([]int, length+1)
	for i := 0; i <= length; i++ {
		cuts[i] = i - 1
	}

	for i := 1; i < length; i++ {
		for j := i; j >= 0; j-- {
			if s[i] == s[j] && ((i-j < 2) || isPa[i-1][j+1]) {
				isPa[i][j] = true
				cuts[i+1] = min(cuts[i+1], cuts[j]+1)
			}
		}
	}
	return cuts[length]
}

func minCut2(s string) int {

	//O(N^2)time O(N^2)space
	length := len(s)
	if length <= 1 {
		return 0
	}

	cuts := make([]int, length+1)
	for i := 0; i <= length; i++ {
		cuts[i] = i - 1
	}

	for i := 0; i < length; i++ {

		for j := 0; i-j >= 0 && i+j < length && s[i-j] == s[i+j]; j++ {
			cuts[i+j+1] = min(cuts[i+j+1], 1+cuts[i-j])
		}
		for j := 1; i-j+1 >= 0 && i+j < length && s[i-j+1] == s[i+j]; j++ {
			cuts[i+j+1] = min(cuts[i+j+1], 1+cuts[i-j+1])
		}
	}
	return cuts[length]
}
