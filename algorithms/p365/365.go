package p365

/**
You are given two jugs with capacities x and y litres. There is an infinite amount of water supply available.
 You need to determine whether it is possible to measure exactly z litres using these two jugs.

If z liters of water is measurable, you must have z liters of water contained within one or both buckets by the end.

Operations allowed:

Fill any of the jugs completely with water.
Empty any of the jugs.
Pour water from one jug into another till the other jug is completely full or the first jug itself is empty.
Example 1: (From the famous "Die Hard" example)

Input: x = 3, y = 5, z = 4
Output: True
Example 2:

Input: x = 2, y = 6, z = 5
Output: False
*/

func gcd(a, b int) int {
	c := 0
	for a != 0 {
		c = a
		a = b % a
		b = c
	}
	return b
}

func canMeasureWater(x int, y int, z int) bool {
	fact := gcd(x, y)
	return z == 0 || z <= (x+y) && z%fact == 0
}
