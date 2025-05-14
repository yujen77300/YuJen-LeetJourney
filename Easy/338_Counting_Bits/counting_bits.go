package countingbits

func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		result[i] = hammingWeight(i)
	}

	return result
}

func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}

	return count
}
