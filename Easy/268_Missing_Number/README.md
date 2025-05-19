# [Missing Number](https://leetcode.com/problems/missing-number/)

## Solution 1: Sorting Approach
- **Time Complexity**: O(n log n) - where n is the length of the array
  - Sorting the array takes O(n log n) time
  - The subsequent linear scan takes O(n) time
- **Space Complexity**: O(1) - constant extra space (if we ignore the space used by sort)
  - We only use a few variables regardless of input size
- **Approach**:
  1. Sort the array in ascending order
  2. Iterate through the sorted array
  3. If the current index doesn't match the value at that index, we've found our missing number
  4. If we complete the iteration without finding a mismatch, the missing number is n (array length)


```go
func missingNumber(nums []int) int {
    sort.Ints(nums)

    for i, v := range nums {
        if i != v {
            return i
        }
    }

    return len(nums)
}

```




## Solution 2: Math Approach (Sum Difference)
- **Time Complexity**: O(n) - where n is the length of the array
  - We iterate through the array exactly once
  - Each operation inside the loop takes constant time
- **Space Complexity**: O(1) - constant extra space
  - Only a single variable (res) regardless of input size
- **Approach**:
  1. Initialize result to n (length of array)
  2. Iterate through indices 0 to n-1
  3. For each index i, add (i - nums[i]) to the result
  4. Return the final result
- **Key Insights**:
  - The expected sum of numbers 0 to n is n*(n+1)/2


```go

func missingNumber(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res += (i - nums[i])
	}
	return res
}

```


