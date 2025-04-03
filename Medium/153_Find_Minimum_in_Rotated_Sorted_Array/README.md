# [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)

## Solution 1: Binary Search
- **Time Complexity**: O(log n) - where n is the length of the array
- **Space Complexity**: O(1) - only using a constant amount of extra space
- **Approach**:
   1. The array can be viewed as two sorted subarrays, with all elements in the second subarray smaller than elements in the first one (unless rotated 0 or n times)
   2. Use binary search to find the pivot point (minimum element)
   3. Compare the middle element with the rightmost element to determine which half contains the minimum
   4. If the middle element is greater than the rightmost element, the minimum is in the right half
   5. If the middle element is less than or equal to the rightmost element, the minimum is in the left half (including the middle element)


```go
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

```