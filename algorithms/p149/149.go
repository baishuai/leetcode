package p149


/**
Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.
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

type slope struct {
	dx int
	dy int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPoints(points []Point) int {
	length := len(points)
	if length <= 2 {
		return length
	}

	result := 0
	for i := 0; i < length; i++ {
		hm := make(map[slope]int)
		samex, samep := 1, 1
		for j := i + 1; j < length; j++ {
			dy, dx := points[j].Y-points[i].Y, points[j].X-points[i].X
			if dx == 0 && dy == 0 {
				samep ++
				continue
			}
			if dx == 0 {
				samex++
				continue
			}
			factor := gcd(dy, dx)
			if factor != 0 {
				dy /= factor
				dx /= factor
			}
			s := slope{dx: dx, dy: dy}
			hm[s]++
		}
		maxp := 0
		for _, v := range hm {
			if v > maxp {
				maxp = v
			}
		}
		maxp += samep

		result = max(result, maxp)
		result = max(result, samex)
	}
	return result
}
