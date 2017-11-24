package p689

func maxSumOfThreeSubarrays(nums []int, k int) []int {

	sum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	leftMax := make([]int, len(nums))
	tot := sum[k]
	cur := 0
	//{7, 13, 20, 19, 19, 2, 10, 1, 1, 19}
	for i := k; i < len(nums); i++ {
		if sum[i]-sum[i-k] > tot {
			tot = sum[i] - sum[i-k]
			cur = i - k
		}
		leftMax[i-1] = cur
	}
	rightMax := make([]int, len(nums))
	tot = sum[len(nums)] - sum[len(nums)-k]
	cur = len(nums) - k
	for j := cur; j > 0; j-- {
		if sum[j+k]-sum[j] > tot {
			tot = sum[j+k] - sum[j]
			cur = j
		}
		rightMax[j] = cur
	}
	res := make([]int, 3)
	max := 0
	for i := k; i <= len(nums)-k*2; i++ {
		m := sum[i+k] - sum[i]
		l := leftMax[i-1]
		m += sum[l+k] - sum[l]
		r := rightMax[i+k]
		m += sum[r+k] - sum[r]

		if m > max {
			max = m
			res[0] = l
			res[1] = i
			res[2] = r
		}
	}
	return res
}
