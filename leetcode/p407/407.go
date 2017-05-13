package p407

/**
Given an m x n matrix of positive integers representing the height of each unit cell in a 2D elevation map, compute the volume of water it is able to trap after raining.

Note:
Both m and n are less than 110. The height of each unit cell is greater than 0 and is less than 20,000.

Example:

Given the following 3x6 height map:
[
  [1,4,3,1,3,2],
  [3,2,1,3,2,4],
  [2,3,3,2,3,1]
]

Return 4.
 */

type Cell struct {
	height int
	x      int
	y      int
}

type MinPq struct {
	pq []Cell
}

func (p *MinPq) swim(k int) {
	for k > 0 && p.pq[(k-1)>>1].height > p.pq[k].height {
		p.pq[(k-1)>>1], p.pq[k] = p.pq[k], p.pq[(k-1)>>1]
		k = (k - 1) >> 1
	}
}

func (p *MinPq) sink(k int) {
	n := len(p.pq)
	for (k<<1)+1 < n {
		j := (k << 1) + 1
		if j < n-1 && p.pq[j].height > p.pq[j+1].height {
			j++
		}
		if p.pq[j].height < p.pq[k].height {
			p.pq[j], p.pq[k] = p.pq[k], p.pq[j]
		} else {
			break
		}
		k = j
	}
}

func (p *MinPq) Insert(v Cell) {
	n := len(p.pq)
	p.pq = append(p.pq, v)
	p.swim(n)
}

func (p *MinPq) Max() Cell {
	return p.pq[0]
}

func (p *MinPq) DelMax() {
	n := len(p.pq)
	p.pq[0], p.pq[n-1] = p.pq[n-1], p.pq[0]
	p.pq = p.pq[:n-1]
	p.sink(0)
}

func (p *MinPq) Empty() bool {
	return len(p.pq) == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m <= 2 {
		return 0
	}
	n := len(heightMap[0])

	minPq := MinPq{pq: make([]Cell, 0)}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		minPq.Insert(Cell{heightMap[i][0], i, 0})
		visited[i][0] = true
		minPq.Insert(Cell{heightMap[i][n-1], i, n - 1})
		visited[i][n-1] = true
	}
	for j := 1; j < n-1; j++ {
		minPq.Insert(Cell{heightMap[0][j], 0, j})
		visited[0][j] = true
		minPq.Insert(Cell{heightMap[m-1][j], m - 1, j})
		visited[m-1][j] = true
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	maxBorder, res := 0, 0
	for !minPq.Empty() {
		c := minPq.Max()
		minPq.DelMax()
		maxBorder = max(maxBorder, c.height)
		for _, d := range dirs {
			x, y := c.x+d[0], c.y+d[1]
			if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] {
				continue
			}
			h := heightMap[x][y]
			if h < maxBorder {
				res += maxBorder - h
				h = maxBorder
			}
			visited[x][y] = true
			minPq.Insert(Cell{h, x, y})
		}
	}
	return res
}
