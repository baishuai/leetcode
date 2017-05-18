package p552

/**
Given a positive integer n, return the number of all possible attendance records with length n, which will be regarded as rewardable. The answer may be very large, return it after mod 109 + 7.

A student attendance record is a string that only contains the following three characters:

'A' : Absent.
'L' : Late.
'P' : Present.
A record is regarded as rewardable if it doesn't contain more than one 'A' (absent) or more than two continuous 'L' (late).

Example 1:
Input: n = 2
Output: 8
Explanation:
There are 8 records with length 2 will be regarded as rewardable:
"PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
Only "AA" won't be regarded as rewardable owing to more than one absent times.

 */

type record struct {
	na_ep  int
	na_el  int
	na_e2l int

	ca_ep  int
	ca_el  int
	ca_e2l int
	ca_ea  int
}

const mod = 1000000000 + 7

func (r *record) all() int {
	return r.na_ep + r.na_el + r.na_e2l +
		r.ca_ep + r.ca_el + r.ca_e2l + r.ca_ea
}

func (r *record) checkMod() {

	r.na_ep %= mod
	r.na_el %= mod
	r.na_e2l %= mod

	r.ca_ep %= mod
	r.ca_el %= mod
	r.ca_e2l %= mod
	r.ca_ea %= mod
}

func plus(rec record) record {
	next := record{}

	next.na_ep = rec.na_ep + rec.na_el + rec.na_e2l
	next.na_el = rec.na_ep
	next.na_e2l = rec.na_el

	next.ca_ep = rec.ca_ep + rec.ca_el + rec.ca_e2l + rec.ca_ea
	next.ca_el = rec.ca_ep + rec.ca_ea
	next.ca_e2l = rec.ca_el
	next.ca_ea = rec.na_ep + rec.na_el + rec.na_e2l

	return next
}

func checkRecord(n int) int {
	if n == 1 {
		return 3
	}

	rec := record{2, 1, 1, 1, 1, 0, 2}

	for i := 3; i <= n; i++ {
		rec = plus(rec)
		rec.checkMod()
	}
	return rec.all() % mod
}
