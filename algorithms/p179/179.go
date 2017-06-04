package p179

import (
	"bytes"
	"sort"
	"strconv"
)

/**
Given a list of non negative integers, arrange them such that they form the largest number.

For example, given [3, 30, 34, 5, 9], the largest formed number is 9534330.

Note: The result may be very large, so you need to return a string instead of an integer.
*/

type Nums []int

func (n Nums) Len() int {
	return len(n)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (n Nums) Less(i, j int) bool {
	i, j = n[i], n[j]
	ipow, jpow := 10, 10
	for (i / ipow) > 0 {
		ipow *= 10
	}
	for (j / jpow) > 0 {
		jpow *= 10
	}
	i, j = i*jpow+j, j*ipow+i
	return i > j
}

// Swap swaps the elements with indexes i and j.
func (n Nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func largestNumber(nums []int) string {
	sort.Sort(Nums(nums))
	var buf bytes.Buffer
	for _, v := range nums {
		buf.WriteString(strconv.Itoa(v))
	}
	bs := bytes.TrimLeft(buf.Bytes(), "0")
	if len(bs) == 0 {
		bs = append(bs, '0')
	}
	return string(bs)
}
