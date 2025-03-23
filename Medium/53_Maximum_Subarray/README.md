# [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

## Solution 1: Two Loops
- **Time Complexity**: O(n²) - where n is the length of the array
- **Space Complexity**: O(1) - only using a constant amount of extra space
- **Approach**:
   1. Use two nested loops to consider every possible subarray
   2. For each starting position, calculate the sum of every possible subarray starting from that position
   3. Keep track of the maximum sum found across all subarrays
   4. <font color=#FF0000>Time Limit Exceeded</font>


```go
func maxSubArray(nums []int) int {
    maxSum := nums[0]

    for i := 0; i < len(nums); i++ {
        currentSum := 0
        for j := i; j < len(nums); j++ {
            currentSum += nums[j]
            if currentSum > maxSum {
                maxSum = currentSum
            }
        }
    }

    return maxSum
}
```


## Solution 2: Kadane's Algorithm
- **Time Complexity**: O(n) - where n is the length of the array
- **Space Complexity**: O(1) - only using a constant amount of extra space
- **Approach**:
  1. For each element, decide whether to start a new subarray or extend the existing one
  2. If current running sum becomes negative, reset it to 0 (equivalent to starting a new subarray)
  3. Keep tracking the maximum sum encountered during the traversal

```go
func maxSubArray(nums []int) int {
	largestSum := nums[0]
	currentSum := 0

	for _, n := range nums {
		if currentSum < 0 {
			currentSum = 0
		}
		currentSum += n
		largestSum = int(math.Max(float64(currentSum), float64(largestSum)))
	}
	return largestSum
}
```