package p042

/**
Given n non-negative integers representing an elevation map where the width of
 each bar is 1, compute how much water it is able to trap after raining.

For example,
Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
 */

func trap(height []int) int {
	rains := 0
	highest, backtrace, waterlevel := 0, 0, 0

	min := func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	for i, v := range height {
		if waterlevel == highest {
			if v >= highest {
				highest = v
				waterlevel = v
				backtrace = i
			}
		} else if waterlevel < highest {
			if v > waterlevel {
				for waterlevel < v && waterlevel < highest {
					if height[backtrace] <= waterlevel {
						backtrace--
						continue
					}
					delta := min(height[backtrace]-waterlevel, v-waterlevel)
					rains += (i - backtrace - 1) * delta
					waterlevel += delta
				}
				if waterlevel == highest {
					highest = v
					backtrace = i
				}
			}
		}
		if v < waterlevel {
			waterlevel = v
			backtrace = i
		}
	}

	return rains
}
