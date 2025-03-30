package binarysearch

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
