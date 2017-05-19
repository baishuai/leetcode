package p282

import "strconv"

/**
Given a string that contains only digits 0-9 and a target value, return all possibilities to add binary operators (not unary) +, -, or * between the digits so they evaluate to the target value.

Examples:
"123", 6 -> ["1+2+3", "1*2*3"]
"232", 8 -> ["2*3+2", "2+3*2"]
"105", 5 -> ["1*0+5","10-5"]
"00", 0 -> ["0+0", "0-0", "0*0"]
"3456237490", 9191 -> []
*/

func addOperators(num string, target int) []string {
	res := make([]string, 0)
	var dfsHelper func(pos int, cv int64, pv int64, cur string)
	dfsHelper = func(pos int, cv int64, pv int64, cur string) {
		if pos == len(num) && cv == int64(target) {
			res = append(res, cur)
		} else {
			for i := pos + 1; i <= len(num); i++ {
				t := num[pos:i]
				now, _ := strconv.ParseInt(t, 10, 0)
				if strconv.FormatInt(now, 10) != t {
					continue
				}
				dfsHelper(i, cv+now, now, cur+"+"+t)
				dfsHelper(i, cv-now, -now, cur+"-"+t)
				dfsHelper(i, cv-pv+pv*now, pv*now, cur+"*"+t)
			}
		}
	}
	if len(num) == 0 {
		return res
	}
	for j := 1; j <= len(num); j++ {
		s := num[0:j]
		cur, _ := strconv.ParseInt(s, 10, 0)
		if strconv.FormatInt(cur, 10) == s {
			dfsHelper(j, cur, cur, s)
		}
	}
	return res
}
