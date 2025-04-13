# [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## Solution 1: Two loops
- **Time Complexity**: O(n log n) - where n is the length of the array
  - Building frequency map: O(n)
  - Creating unique elements slice: O(m) where m is the number of unique elements (worst case m=n)
  - Sorting unique elements by frequency: O(m log m) or O(n log n) in worst case
  - Retrieving top k elements: O(1)
- **Space Complexity**: O(n) - for storing the frequency map and unique elements list
- **Approach**:
  1. Build a hash map to count the frequency of each element
  2. Create a slice of unique numbers from the hash map
  3. Sort the unique numbers by their frequency in descending order
  4. Return the first k elements from the sorted slice

check every single substring

```go
func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }
    longest := s[0:1]

    for i := 0; i < len(s); i++ {
        for j := i + 1; j < len(s); j++ {
            if j-i+1 > len(longest) && isPalindrome(s[i:j+1]) {
                longest = s[i:j+1]
            }
        }
    }

    return longest
}

func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}
```

## Solution 2: Two Pointers To Expand From Center
- **Time Complexity**: O(n²) - where n is the length of the string
  - We consider each position as a potential center: O(n)
  - For each center, we expand outward to find the longest palindrome: O(n)
  - Total: O(n) × O(n) = O(n²)
- **Space Complexity**: O(1) - only constant extra space used regardless of input size
- **Approach**:
  1. Iterate through each position in the string
  2. For each position, expand outward to find palindromes (checking both odd and even length cases)
  3. For odd-length palindromes, start with a single character as center
  4. For even-length palindromes, start with two adjacent characters
  5. Keep track of the longest palindrome found
  6. Return the longest palindromic substring

```go
func longestPalindrome(s string) string {
    result := ""
    maxLength := 0

    for i, _ := range s {
        l := i
        r := i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if (r-l+1) > maxLength{
                result = s[l:r+1]
                maxLength = r-l+1
            }

            l -= 1
            r += 1
        }

        l = i
        r = i + 1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if (r-l+1) > maxLength{
                result = s[l:r+1]
                maxLength = r-l+1
            }

            l -= 1
            r += 1
        }
    } 

    return result
}
```