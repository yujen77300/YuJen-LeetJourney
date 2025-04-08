package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
    l := 0
    r := 0
    answer := 0
    // Use hashMap as a set to check if the character has appeared before
    hashMap := make(map[byte]struct{})

    for r < len(s) {
        if _, isExist := hashMap[s[r]]; !isExist {
            // Add character to the map
            hashMap[s[r]] = struct{}{}
            r += 1
            answer = max(answer, r - l)
        } else {
            // Remove the leftmost character and move left pointer
            delete(hashMap, s[l])
            l += 1
        }
    }

    return answer
}
