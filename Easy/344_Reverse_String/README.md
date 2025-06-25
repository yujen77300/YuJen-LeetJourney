# [Reverse string](https://leetcode.com/problems/reverse-string/)

## Solution 1: Iterative Approach (using temporary variable)
- **Time Complexity**: O(n) - where n is the length of the string
  - Iterate through half of the string, performing constant time operations for each element
- **Space Complexity**: O(1) - constant extra space
  - Only using a temporary variable to store one character during swapping
  - Requires an extra variable for swapping
- **Approach**:
  1. Iterate through the first half of the string
  2. For each position i, swap the character at position i with the character at position (length-i-1)
  3. Use a temporary variable to facilitate the swapping
  4. Continue until we reach the middle of the string


```go
func reverseString(s []byte)  {
    length := len(s)

    for i := 0 ; i < length/2 ; i++{
        temp := s[i]
        s[i] = s[length-i-1]
        s[length-i-1] = temp
    }
}
```


## Solution 2: Two Pointers Approach
- **Time Complexity**: O(n) - where n is the length of the string
  - Visit each character exactly once during the swapping process
- **Space Complexity**: O(1) - constant extra space
  - Only using two pointer variables, no additional data structures needed
- **Approach**:
  1. Initialize two pointers: left at the beginning (0) and right at the end (len(s)-1)
  2. While left pointer is less than right pointer:
     - Swap characters at left and right positions using Go's multiple assignment
     - Move left pointer forward and right pointer backward
  3. Continue until pointers meet in the middle

```go
func reverseString(s []byte) {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}

```