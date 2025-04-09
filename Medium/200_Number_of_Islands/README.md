# [Number of Islands](https://leetcode.com/problems/number-of-islands/)

## Solution 1: DFS
- **Time Complexity**: O(M×N) - where M is the number of rows and N is the number of columns in the grid. Each cell is visited at most once.
- **Space Complexity**: O(M×N) - in the worst case, the entire grid is filled with land, and the recursion stack will have M×N calls.
- **Approach**:
    1. Traverse the entire grid to find unvisited land cells (marked as '1')
    2. When a land cell is found, increment the island count and use DFS to explore the connected land cells
    3. During DFS, mark visited cells as '0' to avoid counting them again
    4. The DFS recursively explores all adjacent land cells in four directions
    5. The total count of DFS initiations gives the number of islands

```go
func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    m, n := len(grid), len(grid[0])
    numIslands := 0

    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
            return
        }

        grid[i][j] = '0'


        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j+1)
        dfs(i, j-1)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                numIslands++
                dfs(i, j)
            }
        }
    }

    return numIslands
}
```