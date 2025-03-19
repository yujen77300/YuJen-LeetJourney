# [Contains Duplicate](https://leetcode.com/problems/contains-duplicate) 

## Solution 1: Set Comparison
- **Time Complexity**: O(n)
- **Space Complexity**: O(n)
- **Approach**: Convert the array to a set to remove duplicates, then compare the lengths. If the set has fewer elements than the original array, duplicates exist.

```python
class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        return (len(set(nums)) != len(nums))
```



## Solution 2: Set Iteration
- **Time Complexity**: O(n)
- **Space Complexity**: O(n)
- **Approach**: Use a set (map in Go) to track seen values. Iterate through the array. If we encounter a value already in the set, return true. Otherwise, add the value to the set and continue.

```golang
func containsDuplicate(nums []int) bool {
    set := make(map[int]struct{})

    for _, num := range nums{
        if _, hasNum := set[num]; hasNum {
            return true
        }
        set[num] = struct{}{}
    }

    return false
}
```