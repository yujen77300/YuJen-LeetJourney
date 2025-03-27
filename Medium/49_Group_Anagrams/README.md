# [Group Anagrams](https://leetcode.com/problems/group-anagrams/)

## Solution 1: Sorting Approach
- **Time Complexity**: O(n * k log k) - where n and k are the lengths of the string array and the maximum length of the string, respectively
- **Space Complexity**: O(n * k) - for storing the map and result
- **Approach**:
   1. Use a hash map to group anagrams together
   2. For each string, sort its characters and use the string as a key
   3. Append the original string to the corresponding group in the map
   4. Convert the map values to a result array



```go
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sortedStr := strings.Join(chars, "")

		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
```