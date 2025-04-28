# [Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/)

## Solution: Hash Map Approach
- **Time Complexity**: O(n) - where n is the length of the array
  - Creating the hash map: O(n)
  - Iterating through unique numbers: O(n)
  - Each number is visited at most twice (once during iteration and possibly once during sequence counting)
- **Space Complexity**: O(n) - for storing all unique numbers in the hash map
- **Approach**:
  1. Create a hash map to store all numbers from the input array
  2. Iterate through the hash map (unique numbers only)
  3. For each number, check if it's the start of a sequence (has no left neighbor)
  4. If it's a sequence start, count consecutive numbers to the right
  5. Keep track of the longest sequence found
  6. Return the length of the longest consecutive sequence

```go
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }

    longest := 0

    // Iterate through each unique number
    for num := range numSet {
        // Only process sequence starts (numbers that have no left neighbor)
        if !numSet[num-1] {
            currentNum := num
            currentStreak := 1

            // Count consecutive numbers to the right
            for numSet[currentNum+1] {
                currentNum++
                currentStreak++
            }

            // Update the longest sequence length
            if currentStreak > longest {
                longest = currentStreak
            }
        }
    }

    return longest
}
```

### Key Insights:
1. By using a hash map, we achieve O(1) lookups to check if numbers exist
2. Only starting counting from sequence start points eliminates redundant work