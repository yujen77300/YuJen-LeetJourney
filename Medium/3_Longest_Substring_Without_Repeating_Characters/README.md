# [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## Solution 1: Sliding Window with Set
- **Time Complexity**: O(2n) ≈ O(n) - where n is the length of the string. In worst case, each character might be visited twice (once by right pointer and once by left pointer).
- **Space Complexity**: O(min(m,n)) - where m is the size of the character set and n is the length of the string. The hashMap stores at most min(m,n) elements.
- **Approach**:
   1. Use two pointers (left and right) to maintain a sliding window
   2. Use a hashMap (as a set) to track characters in the current window
   3. Expand the window to the right if the character is not in the set
   4. Contract the window from the left if a duplicate character is found
   5. Track the maximum window size throughout the process


```go
func lengthOfLongestSubstring(s string) int {
    l := 0
    r := 0
    answer := 0
    hashMap := make(map[byte]struct{})

    for r < len(s) {
        if _, isExist := hashMap[s[r]]; !isExist {
            hashMap[s[r]] = struct{}{}
            r += 1
            answer = max(answer, r - l)
        } else {
            delete(hashMap, s[l])
            l += 1
        }
    }

    return answer
}

```


## Solution 2: Optimized Sliding Window with Index Tracking
- **Time Complexity**: O(n) - where n is the length of the string. Each character is visited exactly once by the right pointer.
- **Space Complexity**: O(min(m,n)) - where m is the size of the character set and n is the length of the string. The hashMap stores at most min(m,n) elements.
- **Approach**:
	1. Use a sliding window technique with a left pointer and right pointer
	2. Use a hashMap to store the index of each character's last occurrence
	3. When encountering a repeat character, jump the left pointer directly to after the last occurrence
	4. Calculate the current window size at each step and track the maximum
	5. This optimization avoids unnecessary iterations of the left pointer

```go
func lengthOfLongestSubstring(s string) int {
    l := 0
    answer := 0
    hashMap := make(map[byte]int)

    for r := 0; r < len(s); r++ {
        if pos, ok := hashMap[s[r]]; ok && pos >= l {
            l = pos + 1
        }

        hashMap[s[r]] = r
        answer = max(answer, r-l+1)
    }

    return answer
}
```
