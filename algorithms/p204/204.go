package p204

/**
Count the number of prime numbers less than a non-negative number, n
*/

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	nums := make([]bool, n)
	nums[0] = true
	nums[1] = true
	cur := 2
	for cur*cur < n && nums[cur] == false {
		tmp := cur * 2
		for tmp < n {
			nums[tmp] = true
			tmp += cur
		}
		cur++
		for nums[cur] && cur*cur < n {
			cur++
		}
	}
	cnt := 0
	for _, v := range nums {
		if v == false {
			cnt++
		}
	}
	return cnt
}
