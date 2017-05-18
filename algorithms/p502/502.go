package p502

/**
Suppose LeetCode will start its IPO soon. In order to sell a good price of its shares to Venture Capital, LeetCode would like to work on some projects to increase its capital before the IPO. Since it has limited resources, it can only finish at most k distinct projects before the IPO. Help LeetCode design the best way to maximize its total capital after finishing at most k distinct projects.

You are given several projects. For each project i, it has a pure profit Pi and a minimum capital of Ci is needed to start the corresponding project. Initially, you have W capital. When you finish a project, you will obtain its pure profit and the profit will be added to your total capital.

To sum up, pick a list of at most k distinct projects from given projects to maximize your final capital, and output your final maximized capital.

Example 1:
Input: k=2, W=0, Profits=[1,2,3], Capital=[0,1,1].

Output: 4

Explanation: Since your initial capital is 0, you can only start the project indexed 0.
             After finishing it you will obtain profit 1 and your capital becomes 1.
             With capital 1, you can either start the project indexed 1 or the project indexed 2.
             Since you can choose at most 2 projects, you need to finish the project indexed 2 to get the maximum capital.
             Therefore, output the final maximized capital, which is 0 + 1 + 3 = 4.
Note:
You may assume all numbers in the input are non-negative integers.
The length of Profits array and Capital array will not exceed 50,000.
The answer is guaranteed to fit in a 32-bit signed integer.
 */

type pAndC struct {
	p int
	c int
}

type greater func(a, b pAndC) bool

type MaxPq struct {
	cmp greater
	pq  []pAndC
}

func (p *MaxPq) swim(k int) {
	for k > 0 && p.cmp(p.pq[k], p.pq[(k-1)>>1]) {
		p.pq[(k-1)>>1], p.pq[k] = p.pq[k], p.pq[(k-1)>>1]
		k = (k - 1) >> 1
	}
}

func (p *MaxPq) sink(k int) {
	n := len(p.pq)
	for (k<<1)+1 < n {
		j := (k << 1) + 1
		if j < n-1 && !p.cmp(p.pq[j], p.pq[j+1]) {
			j++
		}
		if p.cmp(p.pq[j], p.pq[k]) {
			p.pq[j], p.pq[k] = p.pq[k], p.pq[j]
		} else {
			break
		}
		k = j
	}
}

func (p *MaxPq) Insert(v pAndC) {
	n := len(p.pq)
	p.pq = append(p.pq, v)
	p.swim(n)
}

func (p *MaxPq) Max() pAndC {
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

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {

	minC := MaxPq{pq: make([]pAndC, 0), cmp: func(a, b pAndC) bool {
		return a.c < b.c
	}}

	maxP := MaxPq{pq: make([]pAndC, 0), cmp: func(a, b pAndC) bool {
		return a.p > b.p
	}}

	for i := 0; i < len(Profits); i++ {
		minC.Insert(pAndC{p: Profits[i], c: Capital[i]})
	}

	for i := 0; i < k; i++ {
		for !minC.Empty() && minC.Max().c <= W {
			maxP.Insert(minC.Max())
			minC.DelMax()
		}
		if maxP.Empty() {
			break
		} else {
			W += maxP.Max().p
			maxP.DelMax()
		}
	}
	return W
}
