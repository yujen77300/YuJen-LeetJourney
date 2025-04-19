# [Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)

## Solution 1: Greedy Approach
- **Time Complexity**: O(n log n) - where n is the number of intervals
  - Sorting the intervals by start time: O(n log n)
  - Single pass through sorted intervals: O(n)
  - Overall time complexity dominated by sorting: O(n log n)
- **Space Complexity**: O(1) - using only constant extra space regardless of input size
  - Only a few variables are used for tracking results and previous interval end
  - The sort is typically done in-place
- **Approach**:
  1. Sort all intervals based on their start time
  2. Keep track of the end time of the previous non-overlapping interval
  3. Iterate through the sorted intervals:
     - If the current interval starts after or at the previous end, keep it (no overlap)
     - If there's an overlap, increment the removal counter
     - When encountering an overlap, keep the interval with the earlier end time (greedy choice)
  4. Return the total number of intervals that need to be removed

```go
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := 0
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		if start >= prevEnd {
			prevEnd = end
		} else {
			res++
			prevEnd = min(prevEnd, end)
		}
	}

	return res
}
```