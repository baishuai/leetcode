package p480

/**
Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

Examples:
[2,3,4] , the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Your job is to output the median array for each window in the original array.

For example,
Given nums = [1,3,-1,-3,5,3,6,7], and k = 3.

Window position                Median
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7       -1
 1  3 [-1  -3  5] 3  6  7       -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6

Therefore, return the median sliding window as [1,-1,-1,3,5,6].
 */

type MaxPq struct {
	pq []int
}

func (p *MaxPq) swim(k int) {
	for k > 0 && p.pq[(k-1)>>1] < p.pq[k] {
		p.pq[(k-1)>>1], p.pq[k] = p.pq[k], p.pq[(k-1)>>1]
		k = (k - 1) >> 1
	}
}

func (p *MaxPq) sink(k int) {
	n := len(p.pq)
	for (k<<1)+1 < n {
		j := (k << 1) + 1
		if j < n-1 && p.pq[j] < p.pq[j+1] {
			j++
		}
		if p.pq[j] > p.pq[k] {
			p.pq[j], p.pq[k] = p.pq[k], p.pq[j]
		} else {
			break
		}
		k = j
	}
}

func (p *MaxPq) Insert(v int) {
	n := len(p.pq)
	p.pq = append(p.pq, v)
	p.swim(n)
}

func (p *MaxPq) Max() int {
	return p.pq[0]
}

func (p *MaxPq) DelMax() {
	n := len(p.pq)
	p.pq[0], p.pq[n-1] = p.pq[n-1], p.pq[0]
	p.pq = p.pq[:n-1]
	p.sink(0)
}

func (p *MaxPq) Empty() bool {
	return len(p.pq) == 0
}

func medianSlidingWindow(nums []int, k int) []float64 {
	maxHeap := MaxPq{make([]int, 0)}
	minHeap := MaxPq{make([]int, 0)}

	rmHmap := make(map[int]int)
	res := make([]float64, 0)
	ix := 0
	for ix < k {
		maxHeap.Insert(nums[ix])
		ix++
	}
	for i := 0; i < k/2; i++ {
		m := maxHeap.Max()
		maxHeap.DelMax()
		minHeap.Insert(-m)
	}

	pushRes := func() {
		if k%2 == 0 {
			res = append(res, float64(maxHeap.Max()-minHeap.Max())/2)
		} else {
			res = append(res, float64(maxHeap.Max()))
		}
	}

	pushRes()

	for ix < len(nums) {
		n := nums[ix]
		p := nums[ix-k]

		ix ++
		if n == p {
			pushRes()
			continue
		}

		balance := 0
		rmHmap[p] += 1
		if p <= maxHeap.Max() {
			balance--
		} else {
			balance++
		}

		if n <= maxHeap.Max() {
			maxHeap.Insert(n)
			balance++
		} else {
			minHeap.Insert(-n)
			balance--
		}

		if balance < 0 {
			maxHeap.Insert(-minHeap.Max())
			minHeap.DelMax()
		} else if balance > 0 {
			minHeap.Insert(-maxHeap.Max())
			maxHeap.DelMax()
		}

		for !maxHeap.Empty() && rmHmap[maxHeap.Max()] > 0 {
			rmHmap[maxHeap.Max()]--
			if rmHmap[maxHeap.Max()] == 0 {
				delete(rmHmap, maxHeap.Max())
			}
			maxHeap.DelMax()
		}
		for !minHeap.Empty() && rmHmap[-minHeap.Max()] > 0 {
			rmHmap[-minHeap.Max()]--
			if rmHmap[-minHeap.Max()] == 0 {
				delete(rmHmap, -minHeap.Max())
			}
			minHeap.DelMax()
		}
		pushRes()
	}
	return res
}
