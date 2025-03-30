package maximumproductsubarray


func maxProduct(nums []int) int {
    maxEnding := nums[0]
    minEnding := nums[0]
    maxProduct := maxEnding

    for i:=1; i < len(nums) ; i++ {
        temp := minEnding
        minEnding = min(nums[i], nums[i]*minEnding, nums[i]*maxEnding)
        maxEnding = max(nums[i], nums[i]*temp, nums[i]*maxEnding)
        maxProduct = max(maxEnding, maxProduct)
    }

    return maxProduct
}