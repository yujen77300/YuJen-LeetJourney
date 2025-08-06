package maximumproductsubarray

func maxProduct(nums []int) int {
	curMaxProduct := nums[0]
	curMinProduct := nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		temp := curMaxProduct
		curMaxProduct = max(nums[i], curMaxProduct*nums[i], curMinProduct*nums[i])
		curMinProduct = min(nums[i], temp*nums[i], curMinProduct*nums[i])
		maxSoFar = max(maxSoFar, curMaxProduct, curMinProduct)
	}

	return maxSoFar
}
