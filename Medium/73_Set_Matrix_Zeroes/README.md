# [Set Matrix Zeros](https://leetcode.com/problems/set-matrix-zeroes/)

## Solution 1: Store Status In First Col and Row
- **Time Complexity**: O(m × n) - where m is the number of rows and n is the number of columns
  - First pass: O(m × n) to scan the entire matrix and mark zeros
  - Second pass: O(m × n) to set zeros based on markers
  - Additional passes: O(m + n) to handle first row and column
- **Space Complexity**: O(1) - constant extra space
  - Only using two boolean variables to track first row and column status
  - Using the matrix itself as storage for zero markers
- **Approach**:
  1. Check if first row and first column originally contain zeros
  2. Use first row and first column as markers for zero positions
  3. Scan the matrix (excluding first row/column) and mark zeros in first row/column
  4. Use the markers to set zeros in the rest of the matrix
  5. Finally, set first row and column to zero if they originally contained zeros
- **Key Insights**:
  - Must handle first row and column separately to avoid overwriting original data


```go
func setZeroes(matrix [][]int)  {
    rowLength := len(matrix)
    colLength := len(matrix[0])
    firstRowZero := false
    firstColZero := false

    for i:= 0 ; i < colLength ; i++ {
        if matrix[0][i] == 0 {
            firstRowZero = true
            break
        }
    } 

    for j:=0 ; j < rowLength ; j++ {
        if matrix[j][0] == 0 {
            firstColZero = true
            break
        }
    }

    for i := 1 ; i < rowLength ; i++ {
        for j := 1 ; j < colLength ; j ++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    for i := 1; i < rowLength; i++ {
        for j := 1; j < colLength; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }



    if firstRowZero {
        for i := 0; i < colLength; i++ {
            matrix[0][i] = 0
        }
    }

    if firstColZero {
        for i := 0; i < rowLength; i++ {
            matrix[i][0] = 0
        }
    }



}
```