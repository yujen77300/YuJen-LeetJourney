# [Word Break](https://leetcode.com/problems/word-break/)

## Solution: Dynamic Programming
- **Time Complexity**: O(n * m * k) - where n is the length of the string, m is the size of wordDict, and k is the average length of words in wordDict
  - We might need to check each position in the string as a potential starting point
  - For each position, we check against all words in the dictionary
  - Each string comparison takes O(k) time on average
- **Space Complexity**: O(n) - where n is the length of the string
  - We use memoization to store the result for each starting position in the string
  - The recursion call stack can also go as deep as O(n) in the worst case
- **Approach**:
  1. Use memoization to avoid redundant work by storing results for each starting position
  2. Define a recursive function that checks if the string starting from a given index can be segmented
  3. For each position, try all words in the dictionary to see if any word matches the prefix
  4. If a word matches, recursively check if the remainder of the string can be segmented
  5. Cache the result for each starting position to avoid recomputation


```go
func wordBreak(s string, wordDict []string) bool {
    memoization := make(map[int]bool)


    var dp func(startIndex int) bool
    dp = func (startIndex int) bool {
        if startIndex == len(s) {
            return true
        }

        if val, ok := memoization[startIndex] ; ok {
            return val
        }

        for _, word := range wordDict {
            endIndex := startIndex + len(word)

            if endIndex > len(s) {
                continue
            }

            if s[startIndex:endIndex] == word {
                if dp(endIndex) {
                    memoization[startIndex] = true
                    return true
                }
            }
        }

        memoization[startIndex] = false
        return false
    }

    return dp(0)
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


