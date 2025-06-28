package movezeros

func moveZeroes(nums []int) {
	cur := 0

	for _, v := range nums {
		if v != 0 {
			nums[cur] = v
			cur++
		}
	}

	for cur < len(nums) {
		nums[cur] = 0
		// nums = append(nums, 0)
		cur++
	}
}
