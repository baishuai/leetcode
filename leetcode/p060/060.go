package p060

/**
The set [1,2,3,…,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order,
We get the following sequence (ie, for n = 3):

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.
 */

func getPermutation(n int, k int) string {
	bitSet := (1 << uint(n) ) - 1
	kt := make([]int, n-1)
	k--
	fact := 1
	f := 1
	for f < n {
		fact *= f
		f++
	}
	f--
	for i := 0; i < n-1; i++ {
		kt[i] = k / fact
		k = k % fact
		fact /= f
		f--
	}
	ans := make([]byte, n)
	for i := 0; i < n-1; i++ {
		kti := kt[i]
		j := uint(0)
		for kti > 0 {
			if bitSet&(1<<j) != 0 {
				kti--
			}
			j++
		}
		for j < uint(n) {
			if bitSet&(1<<j) != 0 {
				bitSet = bitSet ^ (1 << j)
				ans[i] = byte(j + 1 + '0')
				break
			}
			j++
		}
	}
	for i := uint(0); i < uint(n); i++ {
		if bitSet&(1<<i) != 0 {
			ans[n-1] = byte(i + 1 + '0')
			break
		}
	}
	return string(ans)
}
