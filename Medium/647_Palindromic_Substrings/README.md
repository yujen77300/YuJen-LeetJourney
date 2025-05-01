# [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/)

## Solution 1: Brute Force
- **Time Complexity**: O(n³) - where n is the length of the string
  - We check O(n²) possible substrings
  - Each substring check takes O(n) time to verify if it's palindromic
- **Space Complexity**: O(1) - constant extra space regardless of input size
  - We only use a few variables for tracking indices and counts
- **Approach**:
  1. Generate all possible substrings using two nested loops
  2. For each substring, check if it's a palindrome using two pointers
  3. Count the number of palindromic substrings
  4. Return the final count


```go
func countSubstrings(s string) int {
    count := 0
    for i := 0; i < len(s); i++ {
        for j := i + 1; j <= len(s); j++ {
            if isPalindromic(s[i:j]) {
                count++
            }
        }
    }
    return count
}

func isPalindromic(s string) bool {
    l := 0
    r := len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}

```




## Solution 2: Expand Around Center
- **Time Complexity**: O(n²) - where n is the length of the string
  - We consider n centers (single character) and n-1 centers (between characters)
  - For each center, expansion takes O(n) time in worst case
- **Space Complexity**: O(1) - constant extra space
  - We only use a few variables regardless of input size
  - The closure function doesn't create additional space complexity
- **Approach**:
  1. For each position in the string, expand outward to find all palindromes
  2. Consider two types of centers:
     - Single character center (odd-length palindromes)
     - Between two characters (even-length palindromes)
  3. Count palindromes found during expansion
  4. Return the total count
- **Key Insights**:
  - Use closure functions to encapsulate the expansion logic
  - By expanding around centers, we avoid redundant checks


```go
func countSubstrings(s string) int {
    answer := 0

    expand := func(l, r int) {
        for l >= 0 && r <= len(s) - 1 && s[l] == s[r] {
            answer ++
            l--
            r ++
        }
    }

    for i := 0 ; i <= len(s)-1 ; i++ {
        expand(i,i)
        expand(i,i+1)
    }

    return answer
}
```


