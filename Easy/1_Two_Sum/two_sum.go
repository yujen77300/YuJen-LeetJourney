package twosum

func twoSum(nums []int, target int) []int {
    storeMap := make(map[int]int)

    for index, value := range nums {
        num1 := value
        num2 := target - num1
        if index2,ok := storeMap[num2] ; ok {
            return []int{index,index2}
        }else{
            storeMap[num1] = index
        }
    }
    return []int{}
}