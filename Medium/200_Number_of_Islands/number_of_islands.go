package numberofislands


func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    m, n := len(grid), len(grid[0])
    numIslands := 0

    // DFS : explore the entire island starting from the current position
    var dfs func(i, j int)
    dfs = func(i, j int) {
        // Return if out of bounds or current cell is not land
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
            return
        }

        // Mark current cell as visited by changing it to '0'
        grid[i][j] = '0'

        dfs(i+1, j) // down
        dfs(i-1, j) // up
        dfs(i, j+1) // right
        dfs(i, j-1) // left
    }

    // Traverse the entire grid
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // If we find land ('1'), increment count and use DFS to mark the entire island
            if grid[i][j] == '1' {
                numIslands++
                dfs(i, j)
            }
        }
    }

    return numIslands
}