# [Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## Solution 1: Two Loops
- **Time Complexity**: O(n²) - where n is the length of the array
- **Space Complexity**: O(1) - only using a constant amount of extra space
- **Approach**:
    1. Use two nested loops to consider every possible pair of lines
    2. For each pair, calculate the area by finding minimum height and multiplying by width
    3. Keep track of the maximum area found
    4. Return the maximum area after checking all pairs

```go
func maxArea(height []int) int {
    maxArea := 0
    n := len(height)

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            h := min(height[i], height[j])
            w := j - i
            maxArea = max(maxArea, h * w)
        }
    }

    return maxArea
}
```

## Solution 2: Two Pointers
- **Time Complexity**: O(n) - where n is the length of the array
- **Space Complexity**: O(1) - only using a constant amount of extra space
- **Approach**:
    1. Use two pointers, left and right, starting from the ends of the array
    2. Calculate the area between the two pointers
    3. Move the pointer with the smaller height inward (because moving the taller one would only decrease the area)
    4. Keep track of the maximum area found during this process
    5. Continue until the pointers meet

```go
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
```
