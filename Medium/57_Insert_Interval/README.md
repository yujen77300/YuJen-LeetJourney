# [Insert Interval](https://leetcode.com/problems/insert-interval/)

## Solution 1: Sort and Merge
- **Time Complexity**: O(n log n) - Sort all intervals including the new interval, which takes O(n log n) time. The merging process takes O(n) time.
- **Space Complexity**: O(n) - A new result array to store the merged intervals, which takes O(n) space.
- **Approach**:
  1. Insert the new interval into the existing intervals array
  2. Sort all intervals by their start positions
  3. Initialize a result array and start with the first interval as current
  4. Iterate through remaining intervals and merge overlapping ones
  5. If no overlap exists, add the current interval to result and move to next
- **Key Insights**:
  - The sorting step ensures we process intervals in order, making overlap detection simple
  - We only need to compare adjacent intervals after sorting

```go
func insert(intervals [][]int, newInterval []int) [][]int {
    // 1. Insert newInterval first
    intervals = append(intervals, newInterval)

    sort.Slice(intervals, func(i, j int)bool {
        return intervals[i][0] < intervals[j][0]
    })

    result := [][]int{}
    current := intervals[0]

    // 2. Merge overlapping intervals
    for i := 1 ; i < len(intervals) ;  i ++ {
        if current[1] >= intervals[i][0] {
            if current[1] < intervals[i][1] {
                current[1] = intervals[i][1]
            }
        } else{
            result = append(result, current)
            current = intervals[i]
        }
    }

    result = append(result, current)

    return result
    
}
```


## Solution 2: Linear Scan 
- **Time Complexity**: O(n) - A single pass through the intervals array to find non-overlapping intervals and merge overlapping ones.
- **Space Complexity**: O(n) - A result array to store the final merged intervals.
- **Approach**:
  1. Add all intervals that end before the new interval starts (no overlap on the left)
  2. Merge all intervals that overlap with the new interval by updating the new interval's boundaries
  3. Add the merged new interval to the result
  4. Add all remaining intervals that start after the new interval ends (no overlap on the right)
- **Key Insights**:
  - Since the input intervals are already sorted, we can solve this in linear time without additional sorting

```go
func insert(intervals [][]int, newInterval []int) [][]int {
    result := [][]int{}
    i := 0
    n := len(intervals)

    // 1. Add all intervals that end before newInterval starts (no overlap)
    for i < n && intervals[i][1] < newInterval[0] {
        result = append(result, intervals[i])
        i++
    }

    // 2. Merge all intervals that overlap with newInterval
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    result = append(result, newInterval)

    // 3. Add all remaining intervals that start after newInterval ends
    for i < n {
        result = append(result, intervals[i])
        i++
    }

    return result
}
```