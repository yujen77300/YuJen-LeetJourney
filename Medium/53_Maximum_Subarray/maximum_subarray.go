package maximumsubarray

func maxSubArray(nums []int) int {

	currSum := nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSoFar = max(currSum, maxSoFar)
	}

	return maxSoFar
}
