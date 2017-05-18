package p134

import "testing"

func test(t *testing.T, gas, cost []int, exp int) {
	ans := canCompleteCircuit(gas, cost)
	if ans != exp {
		t.Fail()
	}
}

func TestExample0(t *testing.T) {
	test(t, []int{4}, []int{5}, -1)
}

func TestExtra0(t *testing.T) {
	test(t, []int{6, 0, 1, 3, 2}, []int{4, 5, 2, 5, 5}, -1)
}
