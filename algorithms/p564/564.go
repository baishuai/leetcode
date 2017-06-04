package p564

import (
	"strconv"
)

/**
Given an integer n, find the closest integer (not including itself), which is a palindrome.

The 'closest' is defined as absolute difference minimized between two integers.

Example 1:
Input: "123"
Output: "121"

Note:
The input n is a positive integer represented by string, whose length will not exceed 18.
If there is a tie, return the smaller one as answer.
*/

func reverse(n []byte) []byte {
	res := make([]byte, len(n))
	i, j := 0, len(n)-1
	for i <= j {
		res[i] = n[j]
		res[j] = n[i]
		i++
		j--
	}
	return res
}

func smallerPalindromic(front, mid string) (uint64, string) {
	v, _ := strconv.ParseUint(front+mid, 10, 0)
	small := strconv.FormatUint(v-1, 10)
	var smalls string
	var smallv uint64
	smallbytes := []byte(small)
	if len(small) == len(front+mid) {
		if mid == "" {
			smalls = small + string(reverse(smallbytes))
		} else {
			smalls = small + string(reverse(smallbytes[:len(smallbytes)-1]))
		}
		if smalls == "00" {
			smalls = "9"
		}
	} else {
		if mid != "" {
			smalls = small + string(reverse(smallbytes))
		} else {
			smalls = small + "9" + string(reverse(smallbytes))
		}

	}
	smallv, _ = strconv.ParseUint(smalls, 10, 0)
	return smallv, smalls
}

func largerPalindromic(front, mid string) (uint64, string) {
	v, _ := strconv.ParseUint(front+mid, 10, 0)
	large := strconv.FormatUint(v+1, 10)
	var larges string
	var largev uint64
	largebytes := []byte(large)
	if len(large) == len(front+mid) {
		if mid == "" {
			larges = large + string(reverse(largebytes))
		} else {
			larges = large + string(reverse(largebytes[:len(largebytes)-1]))
		}
	} else {
		if mid != "" {
			larges = large + string(reverse(largebytes[:len(largebytes)-2]))
		} else {
			larges = large + string(reverse(largebytes[:len(largebytes)-1]))
		}
	}
	largev, _ = strconv.ParseUint(larges, 10, 0)
	return largev, larges
}

func nearestPalindromic(n string) string {

	bn := []byte(n)
	length := len(n)
	if length == 1 {
		return string(n[0] - 1)
	}

	origin, _ := strconv.ParseUint(n, 10, 0)

	odd := (length%2 == 1)
	mid := length / 2

	mids := ""
	if odd {
		mids = string(bn[mid])
	}
	front := string(bn[:mid])

	pals := front + mids + string(reverse(bn[:mid]))
	pal, _ := strconv.ParseUint(pals, 10, 0)

	smallerPal, largerPal := uint64(0), uint64(0)
	var smallers, largers string
	if pal < origin {
		smallerPal = pal
		smallers = pals
		largerPal, largers = largerPalindromic(front, mids)
	} else if pal > origin {
		largerPal = pal
		largers = pals
		smallerPal, smallers = smallerPalindromic(front, mids)
	} else {
		largerPal, largers = largerPalindromic(front, mids)
		smallerPal, smallers = smallerPalindromic(front, mids)
	}

	if origin-smallerPal <= largerPal-origin {
		return smallers
	} else {
		return largers
	}
}
