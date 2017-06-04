package p403

/**
A frog is crossing a river. The river is divided into x units and at each unit there may or may not exist a stone. The frog can jump on a stone, but it must not jump into the water.

Given a list of stones' positions (in units) in sorted ascending order, determine if the frog is able to cross the river by landing on the last stone. Initially, the frog is on the first stone and assume the first jump must be 1 unit.

If the frog's last jump was k units, then its next jump must be either k - 1, k, or k + 1 units. Note that the frog can only jump in the forward direction.
*/

func canCross(stones []int) bool {

	if len(stones) == 0 {
		return true
	}
	endStone := stones[len(stones)-1]
	stoneSteps := make(map[int]map[int]struct{})
	var empty struct{}
	stoneSteps[0] = make(map[int]struct{})
	stoneSteps[0][1] = empty

	for i := 1; i < len(stones); i++ {
		stoneSteps[stones[i]] = make(map[int]struct{})
	}

	for _, stone := range stones {
		for step := range stoneSteps[stone] {
			reach := stone + step
			if reach == endStone {
				return true
			}
			if _, ok := stoneSteps[reach]; ok {
				stoneSteps[reach][step] = empty
				stoneSteps[reach][step+1] = empty
				if step-1 > 0 {
					stoneSteps[reach][step-1] = empty
				}
			}
		}

	}
	return false
}
