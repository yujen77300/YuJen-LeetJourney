# [Unique Paths](https://leetcode.com/problems/unique-paths/)

## Solution 1: Top down without memoization (Recursive)
- **Time Complexity**: O(2^(m+n)) - where m and n are the grid dimensions
  - Each cell has two possible directions (down or right)
  - The recursion tree has a depth of m+n-2 (total steps needed)
  - Without memoization, we repeatedly calculate the same subproblems
- **Space Complexity**: O(m+n) - for the recursion stack
  - The maximum depth of recursion is m+n-2
- **Approach**:
  1. Use recursion to calculate the number of unique paths
  2. For each cell, we can reach it from the cell above or the cell to the left
  3. Base case: There is exactly 1 way to reach the starting position (0,0)
  4. Out of bounds positions have 0 paths
  5. For any cell, unique paths = paths from above + paths from left
- **Key Insights**:
  - Time complexity is exponential because we're solving the same subproblems repeatedly

```go
func uniquePaths(m int, n int) int {
	var paths func(i, j int) int
	paths = func(i, j int) int {
		if i == 0 && j == 0 {
			return 1
		} else if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		return paths(i, j-1) + paths(i-1, j)
	}

	return paths(m-1, n-1)
}
```




## Solution 2: Top down with memoziation
- **Time Complexity**: O(m*n) - where m and n are the grid dimensions
  - Each cell's value exactly once
  - There are m*n possible states (i,j combinations)
- **Space Complexity**: O(m*n) - for memoization storage and recursion stack
  - The memo map stores up to m*n entries
  - The recursion stack can grow to m+n depth in worst case
- **Approach**:
  1. Use memoization to avoid recalculating values for the same cell
  2. Store calculated results in a hash map using coordinates as keys
  3. Check if a result is already calculated before making recursive calls
  4. Same base cases as in Solution 1
  5. For any cell, unique paths = paths from above + paths from left
- **Key Insights**:
  - The key optimization is the memo map that reduces time complexity from exponential to polynomial
  - Using array coordinates as map keys allows efficient lookup of previously calculated values


```go
func uniquePaths(m int, n int) int {
	memo := make(map[[2]int]int)
	memo[[2]int{0, 0}] = 1

	var paths func(i, j int) int
	paths = func(i, j int) int {
		if val, ok := memo[[2]int{i, j}]; ok {
			return val
		}
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		val := paths(i-1, j) + paths(i, j-1)
		memo[[2]int{i, j}] = val
		return val
	}

	return paths(m-1, n-1)
}

```
