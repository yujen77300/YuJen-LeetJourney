# [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)

## Solution 1: Sorting Approach
- **Time Complexity**: O(n log n) - where n is the length of the array
  - Building frequency map: O(n)
  - Creating unique elements slice: O(m) where m is the number of unique elements (worst case m=n)
  - Sorting unique elements by frequency: O(m log m) or O(n log n) in worst case
  - Retrieving top k elements: O(1)
- **Space Complexity**: O(n) - for storing the frequency map and unique elements list
- **Approach**:
  1. Build a hash map to count the frequency of each element
  2. Create a slice of unique numbers from the hash map
  3. Sort the unique numbers by their frequency in descending order
  4. Return the first k elements from the sorted slice

```go
func topKFrequent(nums []int, k int) []int {
    countHashMap := make(map[int]int)

    for _, num := range nums {
        countHashMap[num] += 1
    }

    uniqueNums := make([]int, 0, len(countHashMap))
    for num := range countHashMap {
        uniqueNums = append(uniqueNums, num)
    }

    sort.Slice(uniqueNums, func(i, j int) bool {
        return countHashMap[uniqueNums[i]] > countHashMap[uniqueNums[j]]
    })

    return uniqueNums[:k]
}
```

## Solution 2: Bucket Sort Approach
- **Time Complexity**: O(n) - where n is the length of the array
  - Building frequency map: O(n)
  - Creating frequency buckets: O(m) where m is the number of unique elements
  - Collecting results from buckets: O(n) in worst case
- **Space Complexity**: O(n) - for storing the frequency map and buckets
- **Approach**:
  1. Build a hash map to count the frequency of each element
  2. Create frequency buckets where bucket[i] contains all numbers that appear i times
  3. Iterate through the buckets from highest frequency to lowest
  4. Collect the first k elements encountered during the iteration

```go
func topKFrequent(nums []int, k int) []int {
    countHashMap := make(map[int]int)
    for _, num := range nums {
        countHashMap[num] += 1
    }

    buckets := make([][]int, len(nums)+1)
    for num, count := range countHashMap {
        if buckets[count] == nil {
            buckets[count] = []int{}
        }
        buckets[count] = append(buckets[count], num)
    }

    result := []int{}
    for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        if buckets[i] != nil {
            result = append(result, buckets[i]...)
            if len(result) > k {
                result = result[:k]
            }
        }
    }

    return result
}
```