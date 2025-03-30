# [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## Solution 1: Sorting Approach
- **Time Complexity**: O(log n) - where n is the length of the array
- **Space Complexity**: O(1) - using only constant extra space
- **Approach**:
  1. Use a modified binary search algorithm to handle the rotated sorted array
  2. For each mid-point, determine which portion of the array is sorted (left or right)
  3. Based on the sorted portion, check if the target falls within its range
  4. If target is in the sorted portion, search there; otherwise, search in the other portion
  5. Continue binary search until the target is found or the search space is exhausted



```go
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		}

		// left portion
		if nums[l] <= nums[m] {
			if target > nums[m] || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			// right portion
			if target < nums[m] || target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1

}
```