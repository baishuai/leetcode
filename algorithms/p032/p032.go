package p032

/**
Given a string containing just the characters '(' and ')',
 find the length of the longest valid (well-formed) parentheses substring.

For "(()", the longest valid parentheses substring is "()",
 which has length = 2.

Another example is ")()())", where the longest valid parentheses
 substring is "()()", which has length = 4.
*/

func longestValidParentheses(s string) int {
	ss := []byte(s)
	longest := make([]int, len(ss))
	max := 0
	for i, v := range ss {
		if v == ')' {
			if i-1 >= 0 && ss[i-1] == '(' {
				longest[i] = 2
				if i-2 >= 0 {
					longest[i] += longest[i-2]
				}
			} else if i-1 >= 0 && ss[i-1] == ')' {
				if i-1-longest[i-1] >= 0 && ss[i-1-longest[i-1]] == '(' {
					longest[i] = longest[i-1] + 2
					if i-2-longest[i-1] >= 0 {
						longest[i] += longest[i-2-longest[i-1]]
					}

				}
			}
			if longest[i] > max {
				max = longest[i]
			}
		}
	}
	return max
}
