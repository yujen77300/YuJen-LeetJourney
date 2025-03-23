package maximumsubarray

import "math"

func maxSubArray(nums []int) int {
	largestSum := nums[0]
	currentSum := 0

	for _, n := range nums {
		if currentSum < 0 {
			currentSum = 0
		}
		currentSum += n
		largestSum = int(math.Max(float64(currentSum), float64(largestSum)))
	}
	return largestSum
}
