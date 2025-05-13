# [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)

## Solution 1: String Conversion
- **Time Complexity**: O(log n) - where log n represents the number of bits in the number
  - Converting to binary string takes O(log n) time
  - Counting the number of '1's takes O(log n) time
- **Space Complexity**: O(log n) - extra space to store the binary string
  - We need to store a binary string with log n characters
- **Approach**:
  1. Convert the integer to a binary string representation
  2. Count the occurrences of '1' in the binary string
  3. Return the count

```go
func hammingWeight(n int) int {
	binary := strconv.FormatUint(uint64(n), 2)
	return strings.Count(binary, "1")
}

```




## Solution 2: Expand Around Center
- **Time Complexity**: O(log n) - where log n represents the number of bits in the number
- **Space Complexity**: O(1) - constant extra space
- **Approach**:
  1. Initialize a count variable to 0
  2. Iterate until n becomes 0
  3. Check if the least significant bit is 1 using bitwise AND with 1
  4. Increment count if the bit is 1
  5. Right shift n by 1 bit to examine the next bit
  6. Return the final count
- **Key Insights**:
  - Directly manipulates bits without creating any intermediate data structures


```go
func hammingWeight(n int) int {
    count := 0
    for n != 0 {
        if n & 1 == 1 {
            count += 1
        }

        n = n >> 1
    }

    return count
}
```


