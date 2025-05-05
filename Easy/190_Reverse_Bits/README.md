# [Reverse Bits](https://leetcode.com/problems/reverse-bits/)

## Solution: Bit Manipulation
- **Time Complexity**: O(1) - constant time
  - We perform exactly 32 operations regardless of input
  - The bit manipulation operations are all O(1)
- **Space Complexity**: O(1) - constant extra space
  - We only use a few variables regardless of input size
- **Approach**:
  1. Initialize a result variable to 0
  2. Iterate through all 32 bits of the input number
  3. For each bit:
     - Shift result left by 1 bit to make room for the next bit
     - Extract the least significant bit from the input number
     - Add this bit to result using bitwise OR
     - Shift the input number right by 1 bit
  4. Return the final result after processing all 32 bits


```go
func reverseBits(num uint32) uint32 {
    var result uint32 = 0
    for i := 0; i < 32; i++ {
        result = result << 1        // Shift result left by 1 bit to make room
        result |= num & 1           // Extract least significant bit and add to result
        num = num >> 1              // Shift input right by 1 bit to process next bit
    }
    return result
}

```


