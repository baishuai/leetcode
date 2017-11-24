package p388

import (
	"strings"
)

func lengthLongestPath(input string) int {
	lines := strings.Split(input, "\n")

	maxf := 0
	curLen := 0
	flens := make([]int, 0)
	for _, line := range lines {
		l := strings.Count(line, "\t")
		for len(flens) < l {
			flens = append(flens, 0)
		}
		for len(flens) > l {
			curLen -= flens[len(flens)-1]
			flens = flens[0 : len(flens)-1]
		}
		sep := 0
		if l > 0 {
			sep = 1
		}
		flens = append(flens, len(line)-l+sep)
		curLen += flens[len(flens)-1]

		if curLen > maxf && strings.Contains(line, ".") {
			maxf = curLen
		}
	}
	return maxf
}
