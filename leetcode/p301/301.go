package p301

/**
Remove the minimum number of invalid parentheses in order to make the input string valid.
 Return all possible results.

Note: The input string may contain letters other than the parentheses ( and ).

Examples:
"()())()" -> ["()()()", "(())()"]
"(a)())()" -> ["(a)()()", "(a())()"]
")(" -> [""]
 */

const pars = '(' + ')'

type List struct {
	slice []string
}

func reverse(s []byte) []byte {
	res := make([]byte, len(s))
	i, j := 0, len(s)-1
	for i <= j {
		res[i] = s[j]
		res[j] = s[i]
		i++
		j--
	}
	return res
}

func removeInvalidParentheses(s string) []string {
	list := List{slice: make([]string, 0)}
	remove([]byte(s), &list, 0, 0, '(')
	return list.slice
}

func remove(ss []byte, l *List, last_i, last_j int, par byte) {
	for stack, i := 0, last_i; i < len(ss); i++ {
		if ss[i] == par {
			stack++
		} else if ss[i] == pars-par {
			stack--
		}
		if stack >= 0 {
			continue
		}
		for j := last_j; j <= i; j++ {
			if ss[j] == pars-par && (j == last_j || ss[j-1] != pars-par) {
				nss := make([]byte, len(ss)-1)
				copy(nss, ss[:j])
				copy(nss[j:], ss[j+1:])
				remove(nss, l, i, j, par)
			}
		}
		return
	}
	rss := reverse(ss)
	if par == '(' && len(rss) > 0 {
		remove(rss, l, 0, 0, ')')
	} else {
		l.slice = append(l.slice, string(rss))
	}
}
