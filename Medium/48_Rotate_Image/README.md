# [Rotate Image](https://leetcode.com/problems/rotate-image/)

## Solution: In-place Matrix Rotation
- **Time Complexity**: O(n²) - where n is the dimension of the n×n matrix
  - Visit each cell in the matrix once during transpose
  - Visit each cell again during horizontal reflection
- **Space Complexity**: O(1) - constant extra space regardless of input size
  - Only use temporary variables for swapping, no additional data structures
- **Approach**:
  1. Transpose the matrix (convert rows to columns)
  2. Reflect the matrix horizontally (reverse each row)
- **Key Insights**:
  - Rotating a matrix 90° clockwise can be broken down into two simpler operations:
    - Transpose (flip along the diagonal)
    - Horizontal reflection (reverse each row)
  - This approach avoids creating a new matrix, achieving O(1) space complexity

```go
func rotate(matrix [][]int) {
    n := len(matrix)

    // Step 1: Transpose the matrix
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            temp := matrix[i][j]
            matrix[i][j] = matrix[j][i]
            matrix[j][i] = temp
        }
    }

    // Step 2: Reflect horizontally
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            temp := matrix[i][j]
            matrix[i][j] = matrix[i][n-j-1]
            matrix[i][n-j-1] = temp
        }
    }
}
```


