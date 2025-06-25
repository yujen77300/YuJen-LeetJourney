# [Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/)

## Solution 1: Split and Reverse Approach
- **Time Complexity**: O(n) - where n is the length of the string
  - Split operation takes O(n) time to traverse the entire string
  - Reverse iteration takes O(k) time where k is the number of words
  - Join operation takes O(n) time to reconstruct the string
- **Space Complexity**: O(n) - linear space complexity
  - Create an array to store all words after splitting, which takes O(n) space
  - Additional space for the result slice to store non-empty words
- **Approach**:
  1. Split the input string by spaces to get an array of words
  2. Create a new slice to store valid words (non-empty strings)
  3. Iterate through the split array in reverse order
  4. Skip empty strings (caused by multiple consecutive spaces)
  5. Join the valid words with single spaces to form the result

```go
func reverseWords(s string) string {
    allValue := strings.Split(s, " ")
    valueSlice := []string{}

    for i := len(allValue) - 1; i >= 0; i-- {
        if allValue[i] != "" {
            valueSlice = append(valueSlice, allValue[i])
        }
    }

    return strings.Join(valueSlice, " ")
}
```

| Type     | Mutable?  | Notes                         |
| -------- | --------- | ----------------------------- |
| `string` | immutable | Read-only byte sequence       |
| `[]byte` | mutable   | Can modify content            |
| `[]rune` | mutable   | Suitable for Unicode handling |