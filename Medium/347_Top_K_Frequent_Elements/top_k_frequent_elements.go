package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	countHashMap := make(map[int]int)
	for _, num := range nums {
		countHashMap[num] += 1
	}

	// frequency buckets for bucket sort
	buckets := make([][]int, len(nums)+1)
	for num, count := range countHashMap {
		if buckets[count] == nil {
			buckets[count] = []int{}
		}
		buckets[count] = append(buckets[count], num)
	}

	// Collect k elements from high to low frequency
	result := []int{}
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if buckets[i] != nil {
			result = append(result, buckets[i]...)
			if len(result) > k {
				result = result[:k]
			}
		}
	}

	return result
}
