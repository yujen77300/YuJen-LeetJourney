# [Merge Intervals](https://leetcode.com/problems/merge-intervals/)

## Solution 1: Brute 
- **Time Complexity**: O(nlogn) - Other than the sort invocation, we do a simple linear scan of the list, so the runtime is dominated by the O(nlogn) complexity of sorting.
- **Space Complexity**: O(n) - If we can sort intervals in place, we do not need more than constant additional space, although the sorting itself takes O(logn) space. Otherwise, we must allocate linear space to store a copy of intervals and sort that.
- **Approach**:
  1. First sort the intervals by their start positions
  2. Initialize a result array and add the first interval as the current interval being processed
  3. Iterate through the remaining intervals, if the current interval overlaps with the next interval, merge them
  4. If they don't overlap, add the current interval to the result and update the current interval

```go
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    result := [][]int{}
    current := intervals[0]
    for i := 1; i < len(intervals); i++ {
        if current[1] >= intervals[i][0] {
            if intervals[i][1] > current[1] {
                current[1] = intervals[i][1]
            }
        } else {
            result = append(result, current)
            current = intervals[i]
        }
    }

    result = append(result, current)

    return result
}
```
