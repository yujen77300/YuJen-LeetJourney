# [Combination Sum](https://leetcode.com/problems/combination-sum/)

## Solution 1: Backtracking and Index approach
- **Time Complexity**: O(N^(T/M)) where:
  - N is the number of candidates
  - T is the target value
  - M is the minimum value among candidates
  - Since each element can be reused multiple times, the recursion tree can have much more branches than just 2 choices
  - In the worst case, if the smallest candidate is 1, we can have a decision tree of height T with potentially N choices at each level
- **Space Complexity**: O(T/M) - for the recursion stack and current path storage
  - The maximum recursion depth and path length is bounded by T/M
  - In the worst case (when smallest candidate is 1), this becomes O(T)
- **Approach**:
  1. Use backtracking to explore all possible combinations
  2. At each step, we have two choices:
     - Include the current candidate (can be used multiple times)
     - Skip the current candidate and move to the next one
  3. Track the current sum and path during exploration
  4. When current sum equals target, add the current path to results
  5. Use early termination when sum exceeds target or all candidates are processed
- **Key Insights**:
  - This is a combination problem, not permutation, so we use an index parameter to only consider current and future elements
  - We can reuse the same element multiple times, so we stay at index i when including the current element
```go
func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    curPath := []int{}

    var backtracking func (i, curSum int)
    backtracking = func(i, curSum int) {
        if curSum == target {
            // Create a copy of curPath to avoid later modifications affecting our result
            pathCopy := make([]int, len(curPath))
            copy(pathCopy, curPath)
            result = append(result, pathCopy)
            return
        }

        if curSum > target || i == len(candidates) {
            return 
        }

        // Case 1: Include this element (can be used multiple times)
        curPath = append(curPath, candidates[i])
        backtracking(i, curSum + candidates[i])
        // Backtrack: remove the last element
        curPath = curPath[:len(curPath)-1]

        // Case 2: Skip this element, move to the next one
        backtracking(i+1, curSum)
    }

    backtracking(0, 0)
    return result

```