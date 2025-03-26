package mergeintervals

import "sort"

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