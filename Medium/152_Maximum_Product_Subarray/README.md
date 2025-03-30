# [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)

## Solution 1: Two Loops
- **Time Complexity**: O(n²) - where n is the length of the array
- **Space Complexity**: O(1) - only using a constant amount of extra space
- **Approach**:
   1. Use two nested loops to consider every possible subarray
   2. For each starting position, calculate the product of every possible subarray starting from that position
   3. Keep track of the maximum product found across all subarrays


```go
func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    maxProduct := nums[0]
    
    for i := 0; i < len(nums); i++ {
        currentProduct := 1
        for j := i; j < len(nums); j++ {
            currentProduct *= nums[j]
            maxProduct = max(maxProduct, currentProduct)
        }
    }

    return maxProduct
}
```


## Solution 2: Dynamic Programming
- **Time Complexity**: O(n) - where n is the length of the array
- **Space Complexity**: O(1) - only using a constant amount of extra space
- **Approach**:
 1. Unlike maximum sum subarray, we need to track both maximum and minimum products
 2. This is because multiplying a negative number with another negative number can result in a positive product
 3. At each step, we calculate both the maximum and minimum products ending at the current position
 4. We use the current value, or the product of the current value with the previous max/min products
 5. Update the global maximum product as we go through the array

```go
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
```