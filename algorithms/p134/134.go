package p134

/**
There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1).
 You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once, otherwise return -1.

Note:
The solution is guaranteed to be unique.
*/

func canCompleteCircuit(gas []int, cost []int) int {

	if len(gas) != len(cost) {
		return -1
	}
	lens := len(gas)
	tgas := gas[lens-1]
	for i := lens - 1; i > 0; i-- {
		gas[i] = gas[i-1] - cost[i-1]
	}
	gas[0] = tgas - cost[lens-1]

	step := 2 * lens
	start := 0
	index := 0
	countGas := 0
	for step > 1 {
		index = (index + 1) % lens
		countGas += gas[index]
		if countGas < 0 {
			if index == start {
				break
			}
			start = index
			countGas = 0
		} else if index == start {
			return start
		}
		step--
	}
	return -1
}
