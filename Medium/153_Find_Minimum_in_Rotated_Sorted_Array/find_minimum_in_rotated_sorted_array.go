package findminimuminrotatedsortedarray

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]

}
