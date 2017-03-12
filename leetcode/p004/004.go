package p004

/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
 */
import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	//onums1, onums2 := nums1, nums2
	left, right := 0, 0

	for {
		if len(nums2) > len(nums1) {
			nums1, nums2 = nums2, nums1
		}

		//fmt.Printf("%v, %v\n", nums1, nums2)
		m1Index := (len(nums1) - 1) / 2
		m1 := nums1[m1Index]
		l1 := m1Index + 1
		r1 := len(nums1) - m1Index

		m2Index := sort.SearchInts(nums2, m1)
		hit := m2Index < len(nums2) && nums2[m2Index] == m1
		r2 := len(nums2) - m2Index
		l2 := m2Index
		if hit {
			l2++
		}

		diff := (l1 + l2 + left ) - (r1 + r2 + right)
		//fmt.Println(l1, l2, left, r1, r2, right)
		if diff < -1 {
			//去右边找
			if hit {
				l1--
			}
			left += l1 + l2
			nums1 = nums1[l1:]
			nums2 = nums2[l2:]

		} else if diff == -1 {
			if hit {
				return float64(m1)
			}
			var next int
			if r1 > 1 {
				next = nums1[l1]
			}
			if r2 > 0 {
				if (r1 > 1 && next > nums2[l2] ) || r1 <= 1 {
					next = nums2[l2]
				}
			}
			//fmt.Printf("next %v, %v\n", m1, next)
			return float64(m1+next) / 2.0
		} else if diff == 0 {
			return float64(m1)
		} else if diff == 1 {
			if hit {
				return float64(m1)
			}
			var prev int
			if l1 > 1 {
				prev = nums1[l1-2]
			}
			if l2 > 0 {
				if (l1 > 1 && prev < nums2[l2-1]) || l1 <= 1 {
					prev = nums2[l2-1]
				}
			}
			//fmt.Printf("prev %v, %v\n", m1, prev)
			return float64(m1+prev) / 2.0
		} else { // > 1
			right += r1 + r2
			if hit {
				right--
				m1Index++
			}
			nums1 = nums1[:m1Index]
			nums2 = nums2[:m2Index]
		}
	}
}
