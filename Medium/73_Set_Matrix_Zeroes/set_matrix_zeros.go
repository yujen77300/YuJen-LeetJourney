package setmatrixzeroes

func setZeroes(matrix [][]int) {
	rowLength := len(matrix)
	colLength := len(matrix[0])
	firstRowZero := false
	firstColZero := false

	for i := 0; i < colLength; i++ {
		if matrix[0][i] == 0 {
			firstRowZero = true
			break
		}
	}

	for j := 0; j < rowLength; j++ {
		if matrix[j][0] == 0 {
			firstColZero = true
			break
		}
	}

	for i := 1; i < rowLength; i++ {
		for j := 1; j < colLength; j++ {
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
