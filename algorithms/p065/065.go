package p065

/**
Validate if a given string is numeric.

Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
Note: It is intended for the problem statement to be ambiguous.
 You should gather all requirements up front before implementing one.
*/

func isNumber(s string) bool {
	bs := []byte(s)
	l, h := 0, len(bs)-1
	for l <= h && bs[l] == ' ' {
		l++
	}
	for h >= l && bs[h] == ' ' {
		h--
	}
	if l > h {
		return false
	}
	if bs[l] == '-' || bs[l] == '+' {
		l++
	}
	for l <= h && bs[l] >= '0' && bs[l] <= '9' {
		l++
	}
	if l > h {
		return true
	}

	lnum := l > 0 && (bs[l-1] >= '0' && bs[l-1] <= '9')
	rnum := false
	dot := bs[l] == '.'
	if dot {
		l++
		for l <= h && bs[l] >= '0' && bs[l] <= '9' {
			l++
		}
		rnum = bs[l-1] >= '0' && bs[l-1] <= '9'
	}
	if l > h {
		return dot && (lnum || rnum)
	} else if !lnum && !rnum {
		return false
	}

	exp := bs[l] == 'e'
	l++
	if !exp {
		return false
	}
	if l <= h && (bs[l] == '-' || bs[l] == '+') {
		l++
	}
	for l <= h && bs[l] >= '0' && bs[l] <= '9' {
		l++
	}
	if l <= h {
		return false
	}
	rnum = bs[l-1] >= '0' && bs[l-1] <= '9'
	return rnum
}
