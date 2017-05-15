package p483

import (
	"math"
	"strconv"
)

/**
For an integer n, we call k>=2 a good base of n, if all digits of n base k are 1.

Now given a string representing n, you should return the smallest good base of n in string format.

Example 1:
Input: "13"
Output: "3"
Explanation: 13 base 3 is 111.
Example 2:
Input: "4681"
Output: "8"
Explanation: 4681 base 8 is 11111.
Example 3:
Input: "1000000000000000000"
Output: "999999999999999999"
Explanation: 1000000000000000000 base 999999999999999999 is 11.
Note:
The range of n is [3, 10^18].
The string representing n is always valid and will not have leading zeros
 */

func widthBase(n int64, w int) int64 {
	l, r := int64(2), int64(math.Pow(float64(n), 1.0/float64(w-1)))
	res := int64(0)
	for l <= r {
		mean := l + (r-l)>>1
		cnt := int64(0)
		for i := 0; i < w; i++ {
			cnt = cnt*mean + 1
		}
		if cnt < n {
			l = mean + 1 // l-1 is less
		} else if cnt > n {
			r = mean - 1
		} else {
			res = mean
			break
		}
	}
	return res
}

func smallestGoodBase(n string) string {
	nv, _ := strconv.ParseInt(n, 10, 0)
	res := nv - 1
	for w := 64; w >= 2; w-- {
		if base := widthBase(nv, w); base != 0 {
			res = base
			break
		}
	}
	return strconv.FormatInt(res, 10)
}
