# [Valid Anagram](https://leetcode.com/problems/valid-anagram)

## Solution 1: Create two hash map
- **Time Complexity**: O(n) - where n is the length of the strings
- **Space Complexity**: O(k) - where k is the size of the character set (26 for lowercase English letters)
- **Approach**: Create separate hash maps to count occurrences of each character in both strings. First, check if lengths are equal. Then count character frequencies in each string using separate maps. Finally, compare the maps to ensure character counts match exactly.


```go
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	for _, v := range s {
		if time, ok := sMap[v]; ok {
			sMap[v] = time + 1
		} else {
			sMap[v] = 1
		}
	}

	for _, v := range t {
		if time, ok := tMap[v]; ok {
			tMap[v] = time + 1
		} else {
			tMap[v] = 1
		}
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true

}
```


## Solution 2: One map
- **Time Complexity**: O(n) - where n is the length of the strings
- **Space Complexity**: O(k) - where k is the size of the character set (26 for lowercase English letters)
- **Approach**: Use a single hash map for more efficient comparison. First, check if lengths are equal. Then increment counters for each character in the first string. Next, decrement counters for each character in the second string, returning false immediately if any count becomes negative. If both strings are anagrams, all counters will be exactly zero at the end.

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make(map[rune]int)

	for _, v := range s {
		charCount[v]++
	}

	for _, v := range t {
		charCount[v]--
		if charCount[v] < 0 {
			return false
		}
	}

	return true
}
```

```python
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        char_count = {}

        for v in s:
            if v in char_count:
                char_count[v] += 1
            else:
                char_count[v] = 1
        
        for v in t:
            if v not in char_count:
                return False
            char_count[v] -= 1
            if char_count[v] < 0:
                return False
        
        return True
        
```