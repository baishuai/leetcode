package p220

/**
Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.
**/

// 使用滑动窗口（大小为 k）查找，配合二叉树形式的 set 容器，算法复杂度为 O(n * log k）

// 使用 buckets，可以得到线性时间复杂度的算法
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 || k == 0 || t < 0 {
		return false
	}
	buckets := make(map[int]int)
	for i, v := range nums {
		m := v
		if t != 0 {
			m /= t
		}
		if x, ok := buckets[m]; ok && abs(x-v) <= t {
			return true
		} else if l, ok := buckets[m-1]; ok && abs(l-v) <= t {
			return true
		} else if r, ok := buckets[m+1]; ok && abs(r-v) <= t {
			return true
		}
		if i >= k {
			if t != 0 {
				delete(buckets, nums[i-k]/t)
			} else {
				delete(buckets, nums[i-k])
			}
		}
		buckets[m] = v
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
