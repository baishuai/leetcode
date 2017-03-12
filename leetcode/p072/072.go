package p072

/**
Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2.
 (each operation is counted as 1 step.)

You have the following 3 operations permitted on a word:

a) Insert a character
b) Delete a character
c) Replace a character
 */

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func pijDiag(p int, a, b byte) int {
	if a == b {
		return p
	} else {
		return p + 1
	}
}

func minDistance(word1 string, word2 string) int {

	if len(word1) == 0 {
		return len(word2)
	} else if len(word2) == 0 {
		return len(word1)
	}

	dpPre := make([]int, len(word2))
	dpNow := make([]int, len(word2))
	for i := 0; i < len(word1); i++ {
		dpNow, dpPre = dpPre, dpNow
		for j := 0; j < len(word2); j++ {
			if i == 0 && j == 0 {
				dpNow[j] = pijDiag(0, word1[i], word2[j])
			} else if i == 0 {
				dpNow[j] = min(pijDiag(j, word1[i], word2[j]), dpNow[j-1]+1, dpNow[j-1]+1)
			} else if j == 0 {
				dpNow[j] = min(pijDiag(i, word1[i], word2[j]), dpPre[j]+1, dpPre[j]+1)
			} else {
				dpNow[j] = min(pijDiag(dpPre[j-1], word1[i], word2[j]), dpNow[j-1]+1, dpPre[j]+1)
			}
		}
	}
	return dpNow[len(word2)-1]
}
