package palindromicsubstrings

func combinationSum(candidates []int, target int) [][]int {

	result := [][]int{}
	curPath := []int{}

	var backtracking func(i, curSum int)
	backtracking = func(i, curSum int) {
		if curSum == target {
			pathCopy := make([]int, len(curPath))
			copy(pathCopy, curPath)
			result = append(result, pathCopy)
			return
		}

		if curSum > target || i == len(candidates) {
			return
		}

		curPath = append(curPath, candidates[i])
		backtracking(i, curSum+candidates[i])
		curPath = curPath[:len(curPath)-1]

		backtracking(i+1, curSum)
	}

	backtracking(0, 0)
	return result

}
