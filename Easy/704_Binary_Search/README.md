# [Binary Search](https://leetcode.com/problems/binary-search)

## Solution 1:
- **Time Complexity**: O(log n) - where n is the length of the array
- **Space Complexity**: O(1) - using only a constant amount of extra space
- **Approach**:
  1. Initialize two pointers, left and right, pointing to the start and end of the array
  2. Find the middle element between left and right pointers
  3. If the target equals the middle element, return its index
  4. If the target is greater than the middle element, search the right half by moving left pointer
  5. If the target is less than the middle element, search the left half by moving right pointer
  6. Repeat until either the target is found or the search space is exhausted

Go
```go
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
```

Python
```python
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l = 0
        r = len(nums) - 1
        while l <= r :
            m = ( l + r ) // 2
            if target > nums[m]:
                l = m + 1
            elif target < nums[m]:
                r = m - 1
            else:
                return m
        return -1
```