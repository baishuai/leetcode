package p551

/**
You are given a string representing an attendance record for a student.
The record only contains the following three characters:

'A' : Absent.
'L' : Late.
'P' : Present.
A student could be rewarded if his attendance record doesn't contain
more than one 'A' (absent) or more than two continuous 'L' (late).

You need to return whether the student could be rewarded according to his attendance record.

Example 1:
Input: "PPALLP"
Output: True
Example 2:
Input: "PPALLL"
Output: False
*/

func checkRecord(s string) bool {
	l := 0
	a := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			if a {
				return false
			} else {
				a = true
			}
			l = 0
		} else if s[i] == 'L' {
			if l == 2 {
				return false
			} else {
				l += 1
			}
		} else {
			l = 0
		}
	}
	return true
}
