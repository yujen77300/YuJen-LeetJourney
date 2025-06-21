# [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)

## Solution 1: Dynamic Programming
- **Time Complexity**: O(n²) - where n is the length of the input array
  - For each element (O(n)), we compare with all previous elements (O(n))
  - Nested loops result in O(n²) operations
- **Space Complexity**: O(n) - linear extra space
  - Use an array of size n to store the length of LIS ending at each position
  - No additional data structures are needed
- **Approach**:
  1. Initialize an array `result` where `result[i]` represents the length of the longest increasing subsequence ending at index i
  2. Initialize all values in `result` to 1 (minimum LIS length for any position)
  3. For each position i, look at all previous positions j:
     - If current element is greater than previous element (`nums[i] > nums[j]`), we can extend that subsequence
     - Update `result[i]` to maximum of its current value or `result[j] + 1`
  4. Return the maximum value in the result array
- **Key Insights**:
  - Each position maintains its own LIS length, building upon previous results


```go
func lengthOfLIS(nums []int) int {
    n := len(nums)
    result := make([]int,n)

    for i := range result {
        result[i] = 1
    }

    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                result[i] = max(result[i], result[j]+1)
            }
        }
    }

    return slices.Max(result)

}
```


