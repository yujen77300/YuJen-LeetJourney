# [Move Zeroes](https://leetcode.com/problems/move-zeroes/)

## Solution 1: Two-Pass Approach
- **Time Complexity**: O(n) - where n is the length of the array
  - First pass: iterate through the entire array once to move non-zero elements
  - Second pass: fill remaining positions with zeros
  - Total operations are proportional to the array length
- **Space Complexity**: O(1) - constant extra space
  - Only using a single pointer variable (cur) regardless of input size
  - Modify the array in-place without using additional data structures
- **Approach**:
  1. Use a pointer `cur` to track the position for placing non-zero elements
  2. First pass: iterate through the array and move all non-zero elements to the front
     - When encountering a non-zero element, place it at position `cur` and increment `cur`
  3. Second pass: fill all remaining positions from `cur` to the end with zeros
  4. The array is modified in-place maintaining the relative order of non-zero elements
- **Key Insights**:
  - Two-pass approach separates concerns: first collect non-zeros, then fill zeros
  - In-place modification means we cannot use `append()` as it changes slice length


```go
func moveZeroes(nums []int)  {
    cur := 0

    for _, v := range nums {
        if v != 0 {
            nums[cur] = v
            cur ++
        }
    }


    for cur < len(nums) {
        nums[cur] = 0
        // nums = append(nums, 0)
        cur ++
    }
}
```


