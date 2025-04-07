# [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/)

## Solution 1: Brute Force with Two Loops
- **Time Complexity**: O(n) - where n is the length of the string. We process each character at most twice (once when adding to the window, once when removing)
- **Space Complexity**: O(1) - we use a constant-sized map to store character frequencies (at most 26 uppercase English letters)
- **Approach**:
   1. Use a sliding window to track a substring of characters
   2. For each window, find the most frequent character
   3. If (window length - frequency of most common character) ≤ k, the window is valid
   4. Otherwise, shrink the window from the left
   5. Keep track of the maximum valid window size as we go

```go
func characterReplacement(s string, k int) int {
	countMap := make(map[byte]int)
	answer := 0
	maxFreq := 0
	rightIndex := 0

	for leftIndex := range s {
		for rightIndex < len(s) && rightIndex-leftIndex-maxFreq <= k {
			countMap[s[rightIndex]] += 1
			maxFreq = max(maxFreq, countMap[s[rightIndex]])
			rightIndex += 1
		}

		// Update the length of the longest substring
		if rightIndex-leftIndex-maxFreq > k {
			// The new character is different from the current maxFreq character
			answer = max(answer, rightIndex-leftIndex-1)
		} else {
			// rightIndex has reached the rightmost position
			answer = max(answer, rightIndex-leftIndex)
		}

		countMap[s[leftIndex]] -= 1

	}

	return answer

}
```