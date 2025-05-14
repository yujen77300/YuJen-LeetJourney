# [Counting Bits](https://leetcode.com/problems/counting-bits/)

## Solution 1: Bit Counting Approach
- **Time Complexity**: O(n * log n) - where n is the input number
  - We iterate through n+1 numbers (0 to n)
  - For each number, counting bits takes O(log n) time in the worst case
  - Overall complexity is O(n * log n)
- **Space Complexity**: O(n) - for the result array
  - We create an array of size n+1 to store the results
  - Excluding the output array, we use O(1) additional space
- **Approach**:
  1. Create a result array of size n+1
  2. For each number from 0 to n:
     - Count the number of 1 bits using the hammingWeight function
     - Store the count in the result array
  3. Return the result array


```go
func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		result[i] = hammingWeight(i)
	}

	return result
}

func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}

	return count
}
```


## Solution 2: Dynamic Programming
- **Time Complexity**: O(n) - where n is the input number
  - We iterate through n+1 numbers (0 to n)
  - Each calculation takes O(1) time since we use previously computed results
  - Overall complexity is O(n)
  1. Create a result array of size n+1
  2. For each number from 1 to n:
     - Use the relationship: bits in i = bits in (i >> 1) + last bit of i
     - The bits in (i >> 1) are already computed in our DP array
     - The last bit can be obtained using (i & 1)
  3. Return the result array
- **Approach**:
  1. Create a result array of size n+1
  2. For each number from 0 to n:
     - Count the number of 1 bits using the hammingWeight function
     - Store the count in the result array
  3. Return the result array
- **Key Insights**:
  - Remove the last bit of a number (i >> 1) gives a number with a known bit count
  - Dynamic programming reduces time complexity from O(n log n) to O(n)


```go
func countBits(n int) []int {
    result := make([]int, n+1)

    for i := 1; i <= n; i++ {
        result[i] = result[i>>1] + (i & 1)
    }

    return result
}
```