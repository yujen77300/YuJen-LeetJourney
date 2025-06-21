package longestincreasingsubsequence

import "golang.org/x/exp/slices"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	result := make([]int, n)

	for i := range result {
		result[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				result[i] = max(result[i], result[j]+1)
			}
		}
	}

	return slices.Max(result)
}
