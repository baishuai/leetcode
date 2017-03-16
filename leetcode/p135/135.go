package p135

/**
There are N children standing in a line. Each child is assigned a rating value.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
What is the minimum candies you must give?
 */

func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	lastCandy, lastRating := 1, ratings[0]
	candiesCount := 0

	for i := 0; i < len(ratings); i++ {
		v := ratings[i]
		if v > lastRating {
			lastRating = v
			lastCandy ++
			candiesCount += lastCandy
		} else if v == lastRating {
			lastCandy = 1
			candiesCount += lastCandy
		} else {

			highestValue := lastCandy
			j := i
			for j = i + 1; j < len(ratings); j++ {
				if ratings[j] >= ratings[j-1] {
					break
				}
			}
			lastCandy = 1

			candiesCount += ((j - i) + 1 ) * (j - i) / 2

			if (j - i + 1) > highestValue {
				candiesCount += (j - i + 1) - highestValue
			}
			lastCandy = 1
			i = j - 1
			lastRating = ratings[i]
			continue
		}

	}
	return candiesCount
}

func candy_another(ratings []int) int {
	size := len(ratings)
	if size <= 1 {
		return size
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	num := make([]int, size)
	for i := 0; i < size; i++ {
		num[i] = 1
	}
	for i := 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			num[i] = num[i-1] + 1
		}
	}
	for i := size - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			num[i] = max(num[i+1]+1, num[i])
		}
	}
	result := 0
	for i := 0; i < size; i++ {
		result += num[i]
	}
	return result
}
