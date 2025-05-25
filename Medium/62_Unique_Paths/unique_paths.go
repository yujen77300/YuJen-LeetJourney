package uniquepaths

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
