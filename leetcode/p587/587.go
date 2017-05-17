package p587

/**
There are some trees, where each tree is represented by (x,y) coordinate in a two-dimensional garden.
Your job is to fence the entire garden using the minimum length of rope as it is expensive.
The garden is well fenced only if all the trees are enclosed.
Your task is to help find the coordinates of trees which are exactly located on the fence perimeter.

Note:

All trees should be enclosed together. You cannot cut the rope to enclose trees that will separate them in more than one group.
All input integers will range from 0 to 100.
The garden has at least one tree.
All coordinates are distinct.
Input points have NO order. No order required for output.
 */

/**
 * Definition for a point.
 * type Point struct {
 *     X int
 *     Y int
 * }
 */

type Point struct {
	X int
	Y int
}

func outerTrees(points []Point) []Point {
	if len(points) <= 1 {
		return points
	}
	low := points[0]

	for _, p := range points {
		if p.Y < low.Y {
			low = p
		} else if p.Y == low.Y && p.X < low.X {
			low = p
		}
	}

	res := make(map[Point]struct{})
	cur := low
	for {
		res[cur] = struct{}{}

		count := 0
		var pc Point
		tmp := make(map[Point]struct{})
		for _, p := range points {
			if p == cur {
				continue
			}
			count ++
			if count == 1 {
				pc = p
				continue
			}

			if z := (p.X-cur.X)*(pc.Y-cur.Y) - (p.Y-cur.Y)*(pc.X-cur.X); z > 0 {
				pc = p
				tmp = make(map[Point]struct{})
			} else if z == 0 {
				l1 := (p.X-cur.X)*(p.X-cur.X) + (p.Y-cur.Y)*(p.Y-cur.Y)
				l2 := (pc.X-cur.X)*(pc.X-cur.X) + (pc.Y-cur.Y)*(pc.Y-cur.Y)
				tmp[p] = struct{}{}
				tmp[pc] = struct{}{}
				if l1 > l2 {
					pc = p
				}

			}
		}
		cur = pc
		for p := range tmp {
			res[p] = struct{}{}
		}
		if cur == low {
			break
		}
	}
	ans := make([]Point, 0, len(res))
	for p := range res {
		ans = append(ans, p)
	}
	return ans
}
