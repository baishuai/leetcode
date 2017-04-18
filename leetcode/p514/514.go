package p514

import (
	"math"
)

/**
In the video game Fallout 4, the quest "Road to Freedom" requires players to reach a metal dial called the "Freedom Trail Ring", and use the dial to spell a specific keyword in order to open the door.

Given a string ring, which represents the code engraved on the outer ring and another string key, which represents the keyword needs to be spelled. You need to find the minimum number of steps in order to spell all the characters in the keyword.

Initially, the first character of the ring is aligned at 12:00 direction. You need to spell all the characters in the string key one by one by rotating the ring clockwise or anticlockwise to make each character of the string key aligned at 12:00 direction and then by pressing the center button.
At the stage of rotating the ring to spell the key character key[i]:
You can rotate the ring clockwise or anticlockwise one place, which counts as 1 step. The final purpose of the rotation is to align one of the string ring's characters at the 12:00 direction, where this character must equal to the character key[i].
If the character key[i] has been aligned at the 12:00 direction, you need to press the center button to spell, which also counts as 1 step. After the pressing, you could begin to spell the next character in the key (next stage), otherwise, you've finished all the spelling.

 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findRotateSteps(ring string, key string) int {

	positions := make([][]int, 26)
	dps := make([][]int, 26)
	for i := 0; i < 26; i++ {
		positions[i] = make([]int, 0)
	}
	ringLength := len(ring)

	cirleStep := func(i, j int) int {
		if i > j {
			i, j = j, i
		}
		if j-i < (i + ringLength - j) {
			return j - i
		}
		return i + ringLength - j
	}

	for i := 0; i < ringLength; i++ {
		ix := ring[i] - 'a'
		positions[ix] = append(positions[ix], i)
	}
	for i := 0; i < 26; i++ {
		dps[i] = make([]int, len(positions[i]))
	}

	cur := byte(0)
	for i := 0; i < len(key); i++ {
		k := key[i] - 'a'
		if i == 0 {
			for ki, v := range positions[k] {
				dps[k][ki] = cirleStep(0, v)
			}
		} else {
			for ki, pos := range positions[k] {
				tmp := math.MaxInt32
				for ci, prevPos := range positions[cur] {
					tmp = min(tmp, dps[cur][ci]+cirleStep(pos, prevPos))
				}
				dps[k][ki] = tmp
			}
		}
		cur = k
	}
	minStep := math.MaxInt32
	for _, s := range dps[cur] {
		minStep = min(minStep, s)
	}
	return minStep + len(key)
}
