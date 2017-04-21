package p295

/**
Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

Examples:
[2,3,4] , the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:

void addNum(int num) - Add a integer number from the data stream to the data structure.
double findMedian() - Return the median of all elements so far.
For example:

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
 */

// a max heap and a min heap

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

type MedianFinder struct {
	maxHeap MaxPq
	minHeap MaxPq
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: MaxPq{make([]int, 0)},
		minHeap: MaxPq{make([]int, 0)},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Empty() {
		this.maxHeap.Insert(num)
	} else if this.minHeap.Empty() {
		max := this.maxHeap.Max()
		if max <= num {
			this.minHeap.Insert(-num)
		} else {
			this.minHeap.Insert(-max)
			this.maxHeap.DelMax()
			this.maxHeap.Insert(num)
		}
	} else {
		if len(this.maxHeap.pq) == len(this.minHeap.pq) {
			max := this.maxHeap.Max()
			if num <= max {
				this.maxHeap.Insert(num)
			} else {
				this.minHeap.Insert(-num)
			}
		} else if len(this.maxHeap.pq) < len(this.minHeap.pq) {
			min := - this.minHeap.Max()
			if num <= min {
				this.maxHeap.Insert(num)
			} else {
				this.maxHeap.Insert(min)
				this.minHeap.DelMax()
				this.minHeap.Insert(-num)
			}
		} else {
			max := this.maxHeap.Max()
			if num <= max {
				this.minHeap.Insert(-max)
				this.maxHeap.DelMax()
				this.maxHeap.Insert(num)
			} else {
				this.minHeap.Insert(-num)
			}
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	res := 0.0
	if len(this.maxHeap.pq) == len(this.minHeap.pq) {
		res = float64(this.maxHeap.Max()-this.minHeap.Max()) / 2.0
	} else if len(this.maxHeap.pq) < len(this.minHeap.pq) {
		res = float64(-this.minHeap.Max())
	} else {
		res = float64(this.maxHeap.Max())
	}
	return res
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
