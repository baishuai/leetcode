package p187

import "sort"

/**
All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T,
 for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.
Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

For example,
Given s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",

Return:
["AAAAACCCCC", "CCCCCAAAAA"].
*/

type suffixArray [][]byte

// Len is the number of elements in the collection.
func (s suffixArray) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s suffixArray) Less(i, j int) bool {
	minLen := len(s[i])
	if len(s[j]) < minLen {
		minLen = len(s[j])
	}
	for k := 0; k < minLen; k++ {
		if s[i][k] < s[j][k] {
			return true
		} else if s[i][k] > s[j][k] {
			return false
		}
	}
	return len(s[i]) < len(s[j])
}

// Swap swaps the elements with indexes i and j.
func (s suffixArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func dup10(a, b []byte) bool {
	if len(a) < 10 || len(b) < 10 {
		return false
	}
	for i := 0; i < 10; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findRepeatedDnaSequences(s string) []string {
	bs := []byte(s)
	suffixes := make(suffixArray, len(s))
	for i := 0; i < len(s); i++ {
		suffixes[i] = bs[i:]
	}

	sort.Sort(suffixes)
	res := make([]string, 0)
	for i := 0; i < len(s)-1; i++ {
		if dup10(suffixes[i], suffixes[i+1]) {
			s := string(suffixes[i][:10])
			if len(res) == 0 || res[len(res)-1] != s {
				res = append(res, s)
			}
		}
	}
	return res
}
