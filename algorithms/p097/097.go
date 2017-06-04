package p097

/**
Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

For example,
Given:
s1 = "aabcc",
s2 = "dbbca",

When s3 = "aadbbcbcac", return true.
When s3 = "aadbbbaccc", return false.
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	preDp := make([]bool, len(s2)+1)
	nowDp := make([]bool, len(s2)+1)

	bs1, bs2, bs3 := []byte(s1), []byte(s2), []byte(s3)
	for i := 0; i <= len(bs1); i++ {
		preDp, nowDp = nowDp, preDp
		for j := 0; j <= len(bs2); j++ {
			if i == 0 && j == 0 {
				nowDp[j] = true
			} else if i == 0 {
				nowDp[j] = nowDp[j-1] && (bs2[j-1] == bs3[i+j-1])
			} else if j == 0 {
				nowDp[j] = preDp[j] && (bs1[i-1] == bs3[i+j-1])
			} else {
				nowDp[j] = nowDp[j-1] && (bs2[j-1] == bs3[i+j-1]) || preDp[j] && (bs1[i-1] == bs3[i+j-1])
			}
		}
	}
	return nowDp[len(bs2)]
}

/**
type bp struct {
	i int
	j int
	m int
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	// 动态规划 优化加速
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	bs1, bs2, bs3 := []byte(s1), []byte(s2), []byte(s3)
	btrace := []bp{{0, 0, 0}}
	for len(btrace) > 0 {
		bpoint := btrace[len(btrace)-1]
		btrace = btrace[:len(btrace)-1]
		if bpoint.i == len(s1) && bpoint.j == len(s2) {
			return true
		}
		i, j, m := bpoint.i, bpoint.j, bpoint.m
		for ; m < len(s3); m++ {
			match := false
			if i < len(bs1) && bs1[i] == bs3[m] {
				i++
				match = true
			}
			if j < len(bs2) && bs2[j] == bs3[m] {
				if !match {
					j++
				} else {
					btrace = append(btrace, bp{i-1, j + 1, m + 1})
				}
				match = true
			}
			if !match {
				break
			}
		}
		fmt.Println(i, j, m, btrace)
		if i == len(s1) && j == len(s2) {
			return true
		}

	}
	return false
}**/

/** // 递归是正确的，但是非常慢，1778ms
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	return isIntertract([]byte(s1), []byte(s2), []byte(s3))
}
**/

func isIntertract(s1, s2, s3 []byte) bool {
	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	}
	as := false
	if len(s1) > 0 && s1[0] == s3[0] {
		as = as || isIntertract(s1[1:], s2, s3[1:])
	}
	if len(s2) > 0 && s2[0] == s3[0] {
		as = as || isIntertract(s1, s2[1:], s3[1:])
	}
	return as
}
