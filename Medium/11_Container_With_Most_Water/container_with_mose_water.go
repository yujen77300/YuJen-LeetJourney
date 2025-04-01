package containerwithmostwater

func maxArea(height []int) int {
	maxArea := 0

	l := 0
	r := len(height) - 1

	for l != r {
		area := min(height[l], height[r]) * (r - l)
		maxArea = max(maxArea, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
