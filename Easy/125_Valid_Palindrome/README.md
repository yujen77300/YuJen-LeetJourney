# [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/description/)

## Solution: Two Pointers Approach
- **Time Complexity**: O(n) - where n is the length of the string
  - We **traverse the string once** with the two pointers moving toward each other
  - Each character is examined at most once
- **Space Complexity**: O(1) - constant space
  - We only use a fixed amount of variables regardless of input size
  - No additional data structures are used that scale with input size
- **Approach**:
  1. Use two pointers: one starting from the beginning (left) and one from the end (right)
  2. Skip any non-alphanumeric characters by moving the pointers accordingly
  3. Compare characters in a case-insensitive manner
  4. If at any point the characters don't match, return false
  5. If the pointers meet without finding a mismatch, return true

```go
func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l++
		}

		for l < r && !isAlphaNum(s[r]) {
			r--
		}

		if toLower(s[l]) != toLower(s[r]) {
			return false
		}

		l++
		r--
	}

	return true

}

func isAlphaNum(b byte) bool {
	return ('A' <= b && b <= 'Z') ||
		('a' <= b && b <= 'z') ||
		('0' <= b && b <= '9')
}

func toLower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + 32
	}
	return b
}

```
