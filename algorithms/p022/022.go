package p022

/**
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

/**
格路问题
*/

type point struct {
	x int
	y int
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)

	path := make([]byte, n*2)
	for i := 0; i < n; i++ {
		path[i] = '('
	}
	for i := n; i < n*2; i++ {
		path[i] = ')'
	}
	points := make([]point, 2*n-1)
	points[0] = point{n, 0}
	plen := 1

	for {
		//print and append
		ans = append(ans, string(path))
		// no found
		if plen == 2*n-1 {
			break
		}

		// find and then change it
		for i := plen - 1; i >= 0; i-- {
			p := points[i]
			sum := p.x + p.y
			start, index := -1, -1
			if path[sum-1] == '(' {
				if sum >= 2 && path[sum-2] == '(' {
					bp := point{p.x - 1, p.y}
					points[i] = bp
					path[sum-1] = ')'
					secondp := point{bp.x, bp.y + 1}
					points[i+1] = secondp
					thirdp := point{n, secondp.y}
					points[i+2] = thirdp
					start = bp.x
					index = sum
					plen = i + 1 + 2
				}
			} else {
				if p.y < p.x {
					bp := point{p.x, p.y + 1}
					points[i] = bp
					path[sum] = ')'
					secondp := point{n, bp.y}
					points[i+1] = secondp
					start = bp.x
					index = sum + 1
					plen = i + 1 + 1
				}
			}

			if start > 0 && index > 0 {
				for i := start; i < n; i++ {
					path[index] = '('
					index++
				}
				for index < n*2 {
					path[index] = ')'
					index++
				}
				break
			}
		}
	}
	return ans
}
